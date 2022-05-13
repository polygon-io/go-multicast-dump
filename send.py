import socket

MCAST_GRP = '224.0.92.1'
MCAST_PORT = 40001

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, socket.IPPROTO_UDP)
sock.setsockopt(socket.IPPROTO_IP, socket.IP_MULTICAST_TTL, 0)
sock.sendto(("ping from: " + socket.gethostname()).encode('utf-8'), (MCAST_GRP, MCAST_PORT))
