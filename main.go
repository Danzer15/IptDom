import socket
import ipaddress


def execute(user_input):
    try:
        ip = ipaddress.ip_address(user_input)

        domain = socket.gethostbyaddr(str(ip))[0]
        return f"IP {ip} -> Domain: {domain}"
    except ValueError:
        return f"Input '{user_input}' IP not valid"
    except socket.herror:
        return f"IP {user_input} valid, namun tidak memiliki PTR record"
    
if __name__ == "__main__":


    user_input = input("Masukan IP yang valid : ")
    hasil = execute(user_input.strip())
    print(hasil)