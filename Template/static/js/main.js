function showPage(pageId) {
    // 隐藏所有页面
    var pages = document.querySelectorAll('[id^="page"]');
    for (var i = 0; i < pages.length; i++) {
        pages[i].style.display = 'none';
    }

    // 显示选定的页面
    var selectedPage = document.getElementById(pageId);
    if (selectedPage) {
        selectedPage.style.display = 'block';
    }
}

// 获取最新温度和光照数据的函数
function fetchData1() {
    // 发送 GET 请求到后端 API 获取数据
    fetch("http://127.0.0.1:8888/data/check/latestdata")
        .then(response => response.json())
        .then(data => {
            // 更新页面上的温度和光照数据
            document.getElementById("temperature").textContent = data.temperature + "°C";
            document.getElementById("sunlight").textContent = data.sunlit + " lux";
            document.getElementById("personnel").textContent = data.personnel ;
        })
        .catch(error => {
            console.error("Error fetching data:", error);
        });
}
// 初始化页面，获取初始数据,1s更新一次
fetchData1();
setInterval(fetchData1,1000)

// 监听按钮点击事件，点击按钮时更新数据
document.getElementById("updateButton").addEventListener("click", function () {
    fetchData1();// 点击按钮时调用 fetchData 函数更新数据
});
var temperatureButton = document.getElementById("temperatureButton");
var sunlightButton = document.getElementById("sunlightButton");
var peopleButton = document.getElementById("peopleButton");

// 添加点击事件监听器，使用 JavaScript 进行页面跳转
temperatureButton.addEventListener("click", function() {
    window.location.href = "wendu.html"; // 跳转到温度图形页面
});

sunlightButton.addEventListener("click", function() {
    window.location.href = "guangzhao.html"; // 跳转到光照图形页面
});

peopleButton.addEventListener("click", function() {
    window.location.href = "people.html"; // 跳转到光照图形页面
});

//查询历史数据
// var lishi1=document.getElementById("lishi1");
var lishi=document.getElementById("lishi");
// lishi1.addEventListener("click",function (){
//     window.location.href="lishishuju.html";
// });
lishi.addEventListener("click",function (){
    window.location.href="lishishuju.html";
});
// 获取表单和文件输入元素
// var form = document.getElementById("fileUploadForm");
// var fileInput = document.getElementById("fileInput");
// var responseMessage = document.getElementById("responseMessage");
//
// // 监听表单提交事件
// form.addEventListener("submit", function (event) {
//     event.preventDefault(); // 阻止表单默认提交行为
//
//     var formData = new FormData(); // 创建FormData对象
//
//     // 将文件添加到FormData对象中
//     formData.append("file", fileInput.files[0]);
//
//     // 创建一个XMLHttpRequest对象来发送POST请求
//     var xhr = new XMLHttpRequest();
//     xhr.open("POST", "http://127.0.0.1:8888/data/upload/file", true);
//
//     // 监听请求完成事件
//     xhr.onload = function () {
//         if (xhr.status === 200) {
//             alert( "文件上传成功！");
//         } else {
//             alert("文件上传失败：" + xhr.statusText);
//         }
//     };
//
//     // 发送FormData对象
//     xhr.send(formData);
// });
//

// 商品销售
// 获取最新商品数据的函数
function fetchData2() {
    // 发送 GET 请求到后端 API 获取数据
    fetch("http://127.0.0.1:8888/data/check/latestdata")
        .then(response => response.json())
        .then(data => {
            // 更新页面上的商品销售量
            document.getElementById("sell1").textContent = data.a;
            document.getElementById("sell2").textContent = data.b;
            document.getElementById("sell3").textContent = data.c;
            document.getElementById("sell4").textContent = data.d;

            // 获取商品单价元素、销售量元素和销售额元素
            var priceElements = document.querySelectorAll('[id^="price"]');
            var sellElements = document.querySelectorAll('[id^="sell"]');
            var saleElements = document.querySelectorAll('[id^="sale"]');

            // 初始化总销售额
            var totalSales = 0;

            // 遍历商品行，计算每个商品的销售额并更新页面上的销售额元素
            for (var i = 0; i < priceElements.length; i++) {
                var priceElement = priceElements[i];
                var sellElement = sellElements[i];
                var saleElement = saleElements[i];

                // 解析商品单价和销售量的文本内容为浮点数
                var price = parseFloat(priceElement.textContent);
                var sell = parseFloat(sellElement.textContent);

                // 计算销售额
                var sale = price * sell;

                // 更新销售额元素
                saleElement.textContent = "￥" + sale.toFixed(2); // 保留两位小数

                // 更新总销售额
                totalSales += sale;
            }

            // 更新总销售额元素
            var totalSaleElement = document.getElementById("totalSale");
            totalSaleElement.textContent = "￥" + totalSales.toFixed(2); // 保留两位小数
        })
        .catch(error => {
            console.error("Error fetching data:", error);
        });
}
// 初始化页面，获取初始数据,1s更新一次
fetchData2();
setInterval(fetchData2,1000)

// var sell1 = parseFloat(document.getElementById("sell1").textContent);
// var sell2 = parseFloat(document.getElementById("sell2").textContent);
// var sell3 = parseFloat(document.getElementById("sell3").textContent);
// var sell4 = parseFloat(document.getElementById("sell4").textContent);
//
// var sale1 = parseFloat(document.getElementById("sale1").textContent);
// var sale2 = parseFloat(document.getElementById("sale2").textContent);
// var sale3 = parseFloat(document.getElementById("sale3").textContent);
// var sale4 = parseFloat(document.getElementById("sale4").textContent);
//
// var totalSales = (sell1 * sale1) + (sell2 * sale2) + (sell3 * sale3) + (sell4 * sale4);
//
// // 更新销售总额的元素内容
// var totalSaleElement = document.getElementById("totalSale");
// totalSaleElement.textContent = "￥" + totalSales.toFixed(2); // 将总额保留两位小数，并添加货币符号



