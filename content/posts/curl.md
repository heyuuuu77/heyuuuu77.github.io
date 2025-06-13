+++
date = '2025-06-03T14:43:01+08:00'
draft = false
title = '记录一次curl的报错'
+++


```php
    $client_id = env('baidu.api_key');
    $client_secret = env('baidu.secret_key');
    $token_url = 'https://aip.baidubce.com/oauth/2.0/token';
    $curl = curl_init();
    curl_setopt_array($curl, array(
        CURLOPT_URL => "{$token_url}?client_id={$client_id}&client_secret={$client_secret}&grant_type=client_credentials",
        CURLOPT_TIMEOUT => 30,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_CUSTOMREQUEST => 'POST',

        CURLOPT_HTTPHEADER => array(
            'Content-Type: application/json',
            'Accept: application/json'
        ),
    ));
    $response = curl_exec($curl);
    curl_close($curl);

    $response = json_decode($response, 1);
```

$response 始终返回 false, curl_error($curl) 显示的是空字符串 '', curl_errno($curl) 显示的77. 没有其他任何信息。 curl_getinfo($curl) 显示的都是无效信息。 

查询了很多，有的说是 ssl证书 ca文件损坏等。 最终解决方式是重启 php-fpm 服务。 `sudo systemctl restart php-fpm` 就可以了。
