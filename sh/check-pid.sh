while [ True ]
do
	avg_cpu=`top -p 45611 -n 5 -b|grep 45611|awk '{i++}{sum+=$9}END{print sum*100/i}'`
	load=`uptime|awk -F ':|,' '{print $8*100}'`
	echo $avg_cpu, $load
	if [ $avg_cpu -gt 5000 ] || [ $load -gt 500 ] ; then
		echo -e 'start_profile cpu 60\r\n' | nc localhost 13303
		sleep 65
		sudo kill 45611
		sleep 1
		sudo kill 45611
		sleep 1
		sudo kill -9 45611
		sleep 1
		sudo kill -9 45611
		sleep 1
		sudo kill -9 45611
		sleep 1
		echo 'kill -9 45611'
		exit 128
	fi
	sleep 30
done

