#include<stdio.h>
#include<stdlib.h>
#include<string.h>

int split(char* p1, char* p2, char* splits[]) {
    int i = 0;
    int j = 0;
    while(i != -1) {
        int start = i;
        int end = indexOf(p1, p2, start);
        if(end == -1) {
            end = strlen(p1);
        }
        char* p = (char*) malloc(100);
        memcpy(p, p1+start, end-start);
        p[end-start] = '\0';
        splits[j] = p;
        j++;
        i = end + 1;
        if(i > strlen(p1)) {
            break;
        }
    }
    return j;
}
int indexOf(char* p1, char* p2, int start) {
    char* find = strstr(p1 + start, p2);
    if(find != NULL) {
        return find - p1;
    }
    return -1;
}
int is_numberic(char* p1) {
    int len = strlen(p1);
    int i = 0;
    for(i = 0; i < len; i++) {
        char c = p1[i];
        if(c >= '0' && c <= '9') {
            continue;
        }
        return 0;
    }
    return 1;
}
int main(int argc, char **argv) {
    char* splits[100];
    int count = split("20140703160000=action1;20140703160100=action2;1second=action3;weekdays=1,3,5 160000=action4;weekdays=1,3,5 160010=action5", ";", splits);
    printf("count: %d\n", count);
    int i = 0;
    for(i = 0; i < count; i++) {
        char* str = splits[i];
        int index = indexOf(str, "weekdays=", 0);
        if(index == 0) {
            printf("%s: is type3\n", str);
            char* week_splits[2];
            split(str, " ", week_splits);
            char* weekdays = week_splits[0];
            char* actions = week_splits[1];
            printf("weekdays: %s; actions: %s\n", weekdays, actions);
            char* actions_splits[2];
            split(actions, "=", week_splits);
            printf("timer: %s, action: %s\n", week_splits[0], week_splits[1]);
            continue;
        }
        char* sub_splits[2];
        split(str, "=", sub_splits);
        char* key = sub_splits[0];
        char* value = sub_splits[1];
        int numberic = is_numberic(key);
        if(numberic == 1) {
            printf("%s: is type1, key: %s, value: %s\n", str, key, value);
        } else {
            printf("%s: is type2, key: %s, value: %s\n", str, key, value);
        }
        // free sub_splits
    }
    // free splits
}

