<template>
  <html>
    <head>
      <meta charset="UTF-8">
      <title>注册页面</title>
      <style>
        body {
        background-color: #f5f5f5;
        font-family: Arial, sans-serif;
        }

        form {
        margin: 0 auto;
        padding: 20px;
        width: 400px;
        background-color: #fff;
        border: 1px solid #ccc;
        border-radius: 5px;
        }

        h2 {
        text-align: center;
        margin-bottom: 20px;
        }

        label {
        display: block;
        margin-bottom: 5px;
        font-size: 14px;
        }

        input[type="email"], input[type="password"] {
        display: block;
        width: 100%;
        padding: 10px;
        margin-bottom: 10px;
        border-radius: 5px;
        border: 1px solid #ccc;
        font-size: 16px;
        }

        input[type="submit"] {
        display: block;
        width: 100%;
        padding: 10px;
        background-color: #007bff;
        color: #fff;
        border: none;
        border-radius: 5px;
        font-size: 16px;
        cursor: pointer;
        }

        input[type="submit"]:hover {
        background-color: #0069d9;
        }
      </style>
    </head>
    <body>
      <form action="#" method="post">
        <h2>用户注册</h2>
        <label for="email">邮箱：</label>
        <input id="email" type="email" name="email" required>
        <label for="password">密码：</label>
        <input id="password" type="password" name="password" required>
        <label for="confirm_password">确认密码：</label>
        <input id="confirm_password" type="password" name="confirm_password" required>
        <input type="submit" value="提交">
      </form>
    </body>
  </html>

</template>
