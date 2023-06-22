package adapter

// BaseAdapter is a base adapter contains some helper functions.
type BaseAdapter struct {
	db db.Connection
}

// SetConnection set the db connection.
func (base *BaseAdapter) SetConnection(conn db.Connection) {
	base.db = conn
}

// GetConnection get the db connection.
func (base *BaseAdapter) GetConnection() db.Connection {
	return base.db
}

// HTMLContentType return the default content type header.
func (*BaseAdapter) HTMLContentType() string {
	return "text/html; charset=utf-8"
}

// CookieKey return the cookie key.
func (*BaseAdapter) CookieKey() string {
	return auth.DefaultCookieKey
}

// GetUser is a helper function get the auth user model from the context.
func (*BaseAdapter) GetUser(ctx interface{}, wf WebFrameWork) (models.UserModel, bool) {
	cookie, err := wf.SetContext(ctx).GetCookie()

	if err != nil {
		return models.UserModel{}, false
	}

	user, exist := auth.GetCurUser(cookie, wf.GetConnection())
	return user.ReleaseConn(), exist
}

// GetUse is a helper function adds the plugins to the framework.
func (*BaseAdapter) GetUse(app interface{}, plugin []plugins.Plugin, wf WebFrameWork) error {
	if err := wf.SetApp(app); err != nil {
		return err
	}

	for _, plug := range plugin {
		for path, handlers := range plug.GetHandler() {
			if plug.Prefix() == "" {
				wf.AddHandler(path.Method, path.URL, handlers)
			} else {
				wf.AddHandler(path.Method, config.Url("/"+plug.Prefix()+path.URL), handlers)
			}
		}
	}

	return nil
}

func (*BaseAdapter) Run() error         { panic("not implement") }
func (*BaseAdapter) DisableLog()        { panic("not implement") }
func (*BaseAdapter) Static(_, _ string) { panic("not implement") }

// GetContent is a helper function of adapter.Content
func (base *BaseAdapter) GetContent(ctx interface{}, getPanelFn types.GetPanelFn, wf WebFrameWork,
	navButtons types.Buttons, fn context.NodeProcessor) {

	var (
		newBase          = wf.SetContext(ctx)
		cookie, hasError = newBase.GetCookie()
	)

	if hasError != nil || cookie == "" {
		newBase.Redirect()
		return
	}

	user, authSuccess := auth.GetCurUser(cookie, wf.GetConnection())

	if !authSuccess {
		newBase.Redirect()
		return
	}

	var (
		panel types.Panel
		err   error
	)

	if !auth.CheckPermissions(user, newBase.Path(), newBase.Method(), newBase.FormParam()) {
		panel = template.WarningPanel(errors.NoPermission, template.NoPermission403Page)
	} else {
		panel, err = getPanelFn(ctx)
		if err != nil {
			panel = template.WarningPanel(err.Error())
		}
	}

	fn(panel.Callbacks...)

	tmpl, tmplName := template.Default().GetTemplate(newBase.IsPjax())

	buf := new(bytes.Buffer)
	hasError = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(&types.NewPageParam{
		User:         user,
		Menu:         menu.GetGlobalMenu(user, wf.GetConnection(), newBase.Lang()).SetActiveClass(config.URLRemovePrefix(newBase.Path())),
		Panel:        panel.GetContent(config.IsProductionEnvironment()),
		Assets:       template.GetComponentAssetImportHTML(),
		Buttons:      navButtons.CheckPermission(user),
		TmplHeadHTML: template.Default().GetHeadHTML(),
		TmplFootJS:   template.Default().GetFootJS(),
		Iframe:       newBase.Query().Get(constant.IframeKey) == "true",
	}))

	if hasError != nil {
		logger.Error(fmt.Sprintf("error: %s adapter content, ", newBase.Name()), hasError)
	}

	newBase.SetContentType()
	newBase.Write(buf.Bytes())
}