Source code:
<https://github.com/gomatchingdotorg/jsoop>

Ở bài viết này mình sẽ hướng dẫn các bạn áp dụng hướng đối tượng để tính giá sản phẩm lấy từ fakestoreapi.
Bài toán như sau:
Từ 20 sản phẩm lấy từ API, tính tổng giá tổng của sản phẩm:

- Các sản phẩm đều tính thêm 10% thuế VAT
- Sản phẩm thuộc danh mục nam giảm 50% sau thuế
- Sản phẩm thuộc danh mục nữ giảm 30% sau thuế
- Sản phẩm thuộc danh mục trang sức không giảm giá
- Sản phẩm thuộc danh mục đồ điện tử khi mua từ 2 sản phẩm sẻ được giảm 10% một sản phẩm

```js
const CATOGORY_MEN = "men's clothing"
const CATOGORY_WOMEN = "women's clothing"

class Product {
    #id
    #title
    #price
    #category
    constructor(_id,_title, _price,_category) {
        this.#id = _id
this.#title =_title
        this.#price = _price
this.#category =_category
    }
    ToPrice() {
        return this.#price * 1.1
    }
    GetTitle() {
        return this.#title
    }
    GetId() {
        return this.#id
    }
    GetCategory() {
        return this.#category
    }
}

class MenCloth extends Product {
    constructor(_id,_title, _price,_category) {
        super(_id,_title, _price,_category )
    }
    ToPrice() {
        return super.ToPrice() * 0.5
    }
}

class WomenCloth extends Product {
    constructor(_id,_title, _price,_category) {
        super(_id,_title, _price,_category )
    }
    ToPrice() {
        return super.ToPrice() * 0.3
    }
}

let loadData = async function() {
    // fetch('https://fakestoreapi.com/products').then(function(data) {
    //     console.log(data.json())
    // })
    let data = await fetch('https://fakestoreapi.com/products')

    let lstProducts = await data.json()

    console.log(lstProducts) //Array Json

    
    let result = lstProducts.map(function(product) { // Json
        if(product["category"].localeCompare(CATOGORY_MEN) == 0) {
            let sp = new MenCloth(product["id"], product["title"], product["price"], product["category"])
            return sp
        }
        if(product["category"].localeCompare(CATOGORY_WOMEN) == 0) {
            let sp = new WomenCloth(product["id"], product["title"], product["price"], product["category"])
            return sp
        }

        let sp = new Product(product["id"], product["title"], product["price"], product["category"])
        return sp
    })

    console.log(result)  
}
```