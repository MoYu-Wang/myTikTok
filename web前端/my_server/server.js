const express = require('express');
const app = express();
const path = require('path');

// 设置静态文件目录，用于提供 HTML 文件
app.use(express.static(path.join(__dirname, '../public')));

// 启动服务器并监听端口
const PORT = process.env.PORT || 13560;
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
