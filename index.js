import './index.css'
import headerLogoPath from './imgs/big_logo1.png'
import footerLogoPath from './imgs/letters-logo.png'

document.querySelector('.header_logo').src = headerLogoPath
document.querySelector('.footer_logo').src = footerLogoPath


const GetAllProducts = async () => {

    let resp = await fetch('http://localhost:8081/products', {
        method: 'GET',
    });
    let productsMassive = await resp.json()

    productsMassive.forEach(function (product) {
        document.getElementById("products_box").insertAdjacentHTML(`beforeend`,
            `
                 <div class="product">
                   <a class="product_id_link" rel="" >
                     <h2>${product.name}</h2>
                     <img width="100" src="${product.img_url}" alt="">
                     <h4>${product.type}</h4>
                     <h4>Price: ${product.price}</h4>
                   </a>
                 </div>
                 `)
    })
    document.getElementById('products_loading').remove()
    console.log(productsMassive)


}
setTimeout(function () {
    GetAllProducts()
}, 1000);

