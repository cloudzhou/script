for table in `mysql -u git -ppasswd db -e'show tables' |perl -pe 's/\|//g'|grep -v ^Tables_in_`;
do
    mysql -u git -ppasswd db -e"show create table $table";
done

