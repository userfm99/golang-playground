package main

var (
	registrasi = `<!doctype html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta data-n-head="true" name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
 
    <link href="https://fonts.googleapis.com/css?family=Montserrat:400,700&display=swap" rel="stylesheet">
 
    <title>Email Referral</title>

    <style>
        body {
            font-family: 'Montserrat', sans-serif;
            margin: 0 auto;
        }
        .wrapper {
            position: relative;
            width: 100%;
            background-color: #f1f1f1;
            padding-bottom: 25px;
        }
        .wrapper-header, .wrapper-body {
            position: relative;
            display: block;
            padding: 0 20px;
        }
        .wrapper-header {
            padding-top: 20px;
        }
        .wrapper-body {
            padding-bottom: 20px;
        }
        .mail-header {
            position: relative;
            display: block;
            text-align: center;
            background: linear-gradient(#ffc9b3, #f97f4e);
            color: #ffffff;
            padding-top: 30px;
            padding-bottom: 30px;
            border-radius: 8px 8px 0 0;
        }
        .mail-header h3 {
            margin: 0;
        }
        .main-content {
            background-color: #ffffff;
            text-align: center;
            border-radius: 0 0 8px 8px;
            padding-bottom: 50px;
        }
        .img-greeting {
            width: 60%;
            margin: 20px auto;
        }
        @media (min-width: 767px) {
            .wrapper {
                width: 600px;
                padding: 20px;
            }
            .wrapper-header, .wrapper-body {
                padding: 0;
            }
            .img-greeting {
                width: 300px;
            }
        }
    </style>
 
</head>
<body>
    <div class="wrapper">
        <div class="wrapper-header">
            <div class="mail-header">
                <h3>Welcome to Hermales!</h3>
            </div>
        </div>
        <div class="wrapper-body">
            <div class="main-content">
                <img src="https://image.freepik.com/free-vector/group-young-people-posing-photo_52683-18823.jpg" class="img-greeting">
                <div style="padding: 0 15px;">
                    <p style="font-weight: 700; font-size: 18px;text-align: left;">Hallo,</p>
                    <p style="font-size: 14px;color: #6f6f6f;text-align: left;">Selamat datang di hermales.co.id. Anda mendapatkan undangan untuk bergabung. Silahkan lakukan registrasi dengan klik tombol di bawah ini.</p>
                    <p style="font-size: 14px;color: #6f6f6f;text-align: left;margin-bottom: 50px;">Email referral ini hanya berlaku selama 24 jam setelah anda menerima email ini.</p>
                </div>
                <a href="{{link_daftar}}" target="_blank" style="background-color: #627e72;color: #ffffff;font-size: 12px;text-decoration: unset;padding: 10px 15px;border-radius: 4px;">Daftar</a>
            </div>
        </div>
        <p style="text-align: center;font-size: 12px;color: #9e9e9e;">hermales.co.id</p>
    </div>
</body>
</html>
`

	registrasiBerhasil = `<!doctype html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta data-n-head="true" name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
 
    <link href="https://fonts.googleapis.com/css?family=Montserrat:400,700&display=swap" rel="stylesheet">
 
    <title>Email Verification</title>

    <style>
        body {
            font-family: 'Montserrat', sans-serif;
            margin: 0 auto;
        }
        .wrapper {
            position: relative;
            width: 100%;
            background-color: #f1f1f1;
            padding-bottom: 25px;
        }
        .wrapper-header, .wrapper-body {
            position: relative;
            display: block;
            padding: 0 20px;
        }
        .wrapper-header {
            padding-top: 20px;
        }
        .wrapper-body {
            padding-bottom: 20px;
        }
        .mail-header {
            position: relative;
            display: block;
            text-align: center;
            background: linear-gradient(#ffc9b3, #f97f4e);
            color: #ffffff;
            padding-top: 30px;
            padding-bottom: 30px;
            border-radius: 8px 8px 0 0;
        }
        .mail-header h3 {
            margin: 0;
        }
        .main-content {
            background-color: #ffffff;
            text-align: center;
            border-radius: 0 0 8px 8px;
            padding-bottom: 50px;
        }
        .img-greeting {
            width: 60%;
            margin: 20px auto;
        }
        @media (min-width: 767px) {
            .wrapper {
                width: 600px;
                padding: 20px;
            }
            .wrapper-header, .wrapper-body {
                padding: 0;
            }
            .img-greeting {
                width: 300px;
            }
        }
    </style>
 
</head>
<body>
    <div class="wrapper">
        <div class="wrapper-header">
            <div class="mail-header">
                <h3>Hallo, {{username}}</h3>
                <p style="margin-bottom: 0;">Selamat Bergabung di Hermales!</p>
            </div>
        </div>
        <div class="wrapper-body">
            <div class="main-content">
                <img src="https://image.freepik.com/free-vector/group-young-people-posing-photo_52683-18823.jpg" class="img-greeting">
                <div style="padding: 0 15px;">
                    <p style="font-weight: 700; font-size: 18px;text-align: left;">Terimakasih telah mendaftar</p>
                    <p style="font-size: 14px;color: #6f6f6f;text-align: left;">Registrasi Anda telah berhasil. Berikut ini adalah informasi login anda. Gunakan untuk masuk ke Member Area hermales.co.id.</p>
                    <table style="text-align: left;font-size: 14px;color: #6f6f6f;">
                        <tbody>
                            <tr>
                                <td>Username</td>
                                <td>: {{username}}</td>
                            </tr>
                            <tr>
                                <td>Password</td>
                                <td>: {{password}}</td>
                            </tr>
                        </tbody>
                    </table>
                    <p style="font-size: 14px;color: #6f6f6f;text-align: left;margin-bottom: 50px;">Silahkan login ke member area <a href="" target="_blank">disini</a>.</p>
                </div>
            </div>
        </div>
        <p style="text-align: center;font-size: 12px;color: #9e9e9e;">hermales.co.id</p>
    </div>
</body>
</html>`
)
