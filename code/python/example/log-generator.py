# -*- coding: utf-8 -*-
import socket
import json
import time
import random
import datetime

# 配置目标地址
host = 'localhost'
port = 5001

# 创建TCP套接字
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect((host, port))

# 模拟日志级别
levels = ['INFO', 'WARNING', 'ERROR', 'DEBUG']

# 模拟服务名称
services = ['user-service', 'order-service', 'payment-service', 'inventory-service']

# 模拟错误消息
error_messages = [
    'Failed to connect to database',
    'Timeout while waiting for response',
    'Invalid user credentials',
    'Resource not found',
    'Internal server error'
]

try:
    while True:
        # 生成随机日志
        level = random.choice(levels)
        service = random.choice(services)
        message = f"Processing request for {service}"

        # 如果是ERROR级别，使用特殊错误消息
        if level == 'ERROR':
            message = random.choice(error_messages)
        
        # 构建JSON日志
        log_entry = {
            '@timestamp': datetime.datetime.now().isoformat(),
            'level': level,
            'service': service,
            'message': message,
            'request_id': f'{random.randint(100000, 999999)}',
            'duration_ms': random.randint(10, 500)
        }
        
        # 转换为JSON并发送
        log_json = json.dumps(log_entry) + '\n'
        sock.sendall(log_json.encode('utf-8'))
        
        print(f"Sent: {log_json.strip()}")
        
        # 随机间隔
        time.sleep(random.uniform(0.5, 2))
except KeyboardInterrupt:
    print("Stopping log generator...")
finally:
    sock.close()
  