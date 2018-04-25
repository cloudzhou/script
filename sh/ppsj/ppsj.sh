curl http://nation.ppsj.com.cn/jiangshu/index18_list.html|iconv -f gbk -t utf8 > 0.html
for x in {1..188}; do
    curl http://nation.ppsj.com.cn/jiangshu/index18_list115.html|iconv -f gbk -t utf8 > $x.html
    sleep 0.1
done
