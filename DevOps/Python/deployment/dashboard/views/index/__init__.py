from flask import Blueprint, render_template, current_app
from python.modules.checker import get_submodules, add_submodule

import os
static_folder = os.path.join(os.path.dirname(__file__), 'static')

ns = Blueprint(
    'index_blueprint', 
    __name__, 
    template_folder='templates', 
    static_folder=static_folder,
    static_url_path='/index_static'
)


@ns.route('/')
def index():
    return render_template('index.html', modules=get_submodules())