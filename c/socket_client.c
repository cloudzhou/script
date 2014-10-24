#include<stdio.h>
#include<string.h>
#include<sys/socket.h>
#include<arpa/inet.h>
 
int main(int argc, char *argv[]) {
    int sock;
    struct sockaddr_in server;
    char server_reply[2000];
     
    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock == -1) {
        printf("Could not create socket");
        return 1;
    }
     
    server.sin_addr.s_addr = inet_addr("114.215.177.97");
    server.sin_family = AF_INET;
    server.sin_port = htons(8000);
 
    if (connect(sock, (struct sockaddr *)&server, sizeof(server)) < 0) {
        perror("connect failed. Error");
        return 1;
    }
     
    char message[] = "{\"router\": \"FFFFFF01\", \"path\": \"/v1/ping/\", \"method\": \"GET\", \"meta\": {\"Authorization\": \"token 44590e7ccb254203073aeb61777a4c425b20a018\"}}\n";
    if(send(sock, message, strlen(message), 0) < 0) {
        puts("Send failed");
        return 1;
    }
         
    if(recv(sock, server_reply, 2000, 0) < 0) {
        puts("recv failed");
        return 1;
    }
         
    printf("get response: %s", server_reply);
    close(sock);
    return 0;
}
