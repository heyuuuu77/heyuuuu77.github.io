+++
date = '2025-01-10T10:36:43+08:00'
draft = true
title = 'Hey'
tags = ["心理"]
+++


hey
<!--more-->


<style>
    /* 音频播放器容器样式 */
    .audio-player {
        width: 300px;
        background-color: #f4f4f4;
        border-radius: 10px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        overflow: hidden;
        margin: 20px auto;
    }

    /* 封面图片样式 */
    .audio-cover {
        width: 100%;
        height: 200px;
        background-image: url('/images/tabudong.jpg');
        background-size: cover;
        background-position: center;
    }

    /* 音频控件样式 */
    audio {
        width: 100%;
        margin-top: -4px;
    }
</style>

<div class="audio-player">
    <div class="audio-cover"></div>
    <audio controls>
        <source src="/audio/他不懂.mp3" type="audio/mpeg">
        您的浏览器不支持音频播放。
    </audio>
</div>

