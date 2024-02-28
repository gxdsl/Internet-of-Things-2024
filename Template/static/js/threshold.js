//阈值
var thresholdForm = document.getElementById("up_threshold");

// 添加表单提交事件监听器
thresholdForm.addEventListener("submit", function(event) {
    event.preventDefault(); // 阻止默认表单提交行为

    var t_h = document.getElementById("temh").value;
    var t_l = document.getElementById("teml").value;
    var s_h = document.getElementById("lighth").value;
    var s_l = document.getElementById("lightl").value;
    // 创建一个FormData对象，用于将表单数据发送给后端
    var formData = new FormData();
    formData.append("temh", t_h);
    formData.append("teml", t_l);
    formData.append("lighth",s_h);
    formData.append("lightl",s_l);
    fetch("http://127.0.0.1:8888/status/updatedata", {
        method: "POST",
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            // 处理后端响应
            if (data.code === 200) {
                alert(data.message)
            } else {
                alert(data.message);
            }
        })
});

// 获取阈值数据的函数
function fetch_the() {
    // 发送 GET 请求到后端 API 获取数据
    fetch("http://127.0.0.1:8888/status/getdata")
        .then(response => response.json())
        .then(data => {
            // 更新页面数据
            document.getElementById("tem_h").textContent = data.temh + "°C";
            document.getElementById("tem_l").textContent = data.teml + "°C";
            document.getElementById("light_h").textContent = data.lighth + " lux";
            document.getElementById("light_l").textContent = data.lightl + " lux";
        })
        .catch(error => {
            console.error("Error fetching data:", error);
        });

}
// 初始化页面，获取初始数据
fetch_the();
setInterval(fetch_the,1000);