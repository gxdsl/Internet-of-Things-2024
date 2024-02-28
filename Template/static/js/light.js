const toggleButton = document.getElementById("toggleButton");
const lamp = document.getElementById("lamp");

// 发送GET请求来获取初始灯状态
getLampStatus();
setInterval(getLampStatus, 1000);
toggleButton.addEventListener("click", function () {
    // 发送POST请求到后端API，切换灯状态
    fetch("http://127.0.0.1:8888/status/togglelamp", {
        method: "POST"
    })
        .then(response => response.json())
        .then(data => {
            // 根据后端返回的状态更新前端灯泡状态
            if (data.lamp === 1) {
                // 灯开启，添加 "on" 类，移除 "off" 类
                lamp.classList.add("on");
                lamp.classList.remove("off");
            } else {
                // 灯关闭，添加 "off" 类，移除 "on" 类
                lamp.classList.add("off");
                lamp.classList.remove("on");
            }
        });
});

// 获取灯状态的函数
function getLampStatus() {
    // 发送GET请求到后端API，获取灯状态
    fetch("http://127.0.0.1:8888/status/getlamp")
        .then(response => response.json())
        .then(data => {
            // 根据后端返回的状态更新前端灯泡状态
            if (data.lamp === 1) {
                // 灯开启，添加 "on" 类，移除 "off" 类
                lamp.classList.add("on");
                lamp.classList.remove("off");
            } else {
                // 灯关闭，添加 "off" 类，移除 "on" 类
                lamp.classList.add("off");
                lamp.classList.remove("on");
            }
        });
}

