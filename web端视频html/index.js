const { app, BrowserWindow} = require('electron');
const path = require('path');

function createWindow() {
    const mainWindow = new BrowserWindow({
        width: 1200,
        height: 850,
        autoHideMenuBar: true, // 设置为 true 隐藏菜单
        webPreferences: {
            nodeIntegration: true
        }
    });

    //处理聚焦问题
    const isWindows = process.platform === 'win32';
    let needsFocusFix = false;
    let triggeringProgrammaticBlur = false;
    mainWindow.on('blur', (event) => {
      if(!triggeringProgrammaticBlur) {
        needsFocusFix = true;
      }
    })
    mainWindow.on('focus', (event) => {
      if(isWindows && needsFocusFix) {
        needsFocusFix = false;
        triggeringProgrammaticBlur = true;
        setTimeout(function () {
            mainWindow.blur();
            mainWindow.focus();
          setTimeout(function () {
            triggeringProgrammaticBlur = false;
          }, 100);
        }, 100);
      }
    })

    // 指定你的 HTML 文件
    mainWindow.loadFile(path.join(__dirname, 'login.html')); 
}

app.on('ready', createWindow);
