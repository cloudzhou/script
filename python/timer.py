# rule: secnod(0-59) minute(0-59) hour(0-23) day(1-31) month(1-12) year(yyyy) day_of_week(0-7, 0 or 7 is Sun) period(20140623101010-20140723101010)
# for every field: 1) * 2) */5 3) 1,3,5,7 4) 1-3,7-9 5) 2-7/2 
import re
import time
from datetime import datetime

rules = [
'7 38 16 2 7 2014 * * fixtime',
'*/9 * * * * * * * every_15seconds',
'* */2 * * * * * * every_minute',
'* * * * * * * 180458,185317 time_period_in_day_start,time_period_in_day_end',
'* * * * * * 3 175257,175317 time_period_in_day_start_week,time_period_in_day_end_week',
'* * * * * * 4,5 175257,175317 time_period_in_day_start_week2,time_period_in_day_end_week2',
'* * * * * * * 20140702171210,20140702171217 time_period_start,time_period_end',
]
'*/15 * * * * * * 20140702101010-20140702112610 time_period_every_15seconds',
'* */1 * * * * 1,3 20140702101010-20140702112610 time_period_every_1minute_monday_wednesday',
def main():
    print 'start timer'
    while True:
        now = datetime.now()
        seconds_since_epoch = int(datetime.now().strftime("%s"))
        (year, month, day, hour, minute, second, weekday) = (now.year, now.month, now.day, now.hour, now.minute, now.second, now.isoweekday())
        min_wait_seconds = -1
        run_actions = ''
        for rule in rules:
            (r_second, r_minute, r_hour, r_day, r_month, r_year, r_weekday, r_period, r_action) = rule.split(' ')
            if re.match('\d+ \d+ \d+ \d+ \d+ \d+ \* \* \w+', rule):
                #print '1 at fix time:'
                #print '    ' + rule
                fixtime = datetime.strptime('%s-%s-%s %s:%s:%s' % (r_year, r_month, r_day, r_hour, r_minute, r_second), '%Y-%m-%d %H:%M:%S')
                #print fixtime
                wait_seconds = get_wait_seconds(fixtime, now)
                if min_wait_seconds == -1 or wait_seconds <= min_wait_seconds:
                    if wait_seconds == min_wait_seconds:
                        run_actions = run_actions + ' ' + r_action
                    else:
                        run_actions = r_action
                    min_wait_seconds = wait_seconds
            elif re.match('\* \* \* \* \* \* \* (\d{14},)+\d{14} (\w+,)+\w+', rule):
                #print '2 at fix multi time:'
                #print '    ' + rule
                for fixtime_str in r_period.split(','):
                    fixtime = datetime.strptime(fixtime_str, '%Y%m%d%H%M%S')
                    #print fixtime
                    wait_seconds = get_wait_seconds(fixtime, now)
                    if min_wait_seconds == -1 or wait_seconds <= min_wait_seconds:
                        if wait_seconds == min_wait_seconds:
                            run_actions = run_actions + ' ' + r_action
                        else:
                            run_actions = r_action
                        min_wait_seconds = wait_seconds
            m = re.match('.*\*/(\d+).*', rule)
            if m:
                #print '3 at loop time:'
                #print '    ' + rule
                loop_period = int(m.group(1))
                wait_seconds = None
                if re.match('\*/\d+', r_second):
                    wait_seconds = loop_period - (seconds_since_epoch % loop_period)
                    #print 'aaaaaa %s %s %s' % (wait_seconds, seconds_since_epoch, loop_period)
                elif re.match('\*/\d+', r_minute):
                    wait_seconds = loop_period*60 - (seconds_since_epoch % (loop_period*60))
                elif re.match('\*/\d+', r_hour):
                    wait_seconds = loop_period*60*60 - (seconds_since_epoch % (loop_period*60*60))
                elif re.match('\*/\d+', r_day):
                    wait_seconds = loop_period*60*60*24 - (seconds_since_epoch % (loop_period*60*60*24))
                elif re.match('\*/\d+', r_month):
                    wait_seconds = loop_period*60*60*24*30 - (seconds_since_epoch % (loop_period*60*60*24*30))
                elif re.match('\*/\d+', r_year):
                    wait_seconds = loop_period*60*60*24*30*365 - (seconds_since_epoch % (loop_period*60*60*24*30*365))
    
                if wait_seconds is not None:
                    if min_wait_seconds == -1 or wait_seconds <= min_wait_seconds:
                        if wait_seconds == min_wait_seconds:
                            run_actions = run_actions + ' ' + r_action
                        else:
                            run_actions = r_action
                        min_wait_seconds = wait_seconds
            elif re.match('\* \* \* \* \* \* \* (\d{6},)+\d{6} (\w+,)+\w+', rule) or re.match('\* \* \* \* \* \* (\d+,)+\d+ (\d{6},)+\d{6} (\w+,)+\w+', rule):
                #print '4 at loop period time(every day):'
                #print '    ' + rule
                weekdays = [1, 2, 3, 4, 5, 6, 7]
                if r_weekday != '*':
                    weekdays = r_weekday.split(',')
                    weekdays.sort()
                days = 0
                i = 0
                while i < len(weekdays):
                    _weekdays = int(weekdays[i])
                    if weekday <= _weekdays:
                        days = _weekdays - weekday
                        break
                    if i == len(weekdays) - 1:
                        days = 7 - weekday + int(weekdays[0])
                        break
                    i = i + 1
                print 'days: %s' % days
                for hour_minute_second in r_period.split(','):
                    (_hour, _minute, _second) = (hour_minute_second[0:2], hour_minute_second[2:4], hour_minute_second[4:6])
                    wait_seconds = int(days)*3600 + int(_hour)*3600 + int(_minute)*60 + int(_second) - hour*3600 - minute*60 - second
                    if wait_seconds > 0 and wait_seconds <= min_wait_seconds:
                        if wait_seconds == min_wait_seconds:
                            run_actions = run_actions + ' ' + r_action
                        else:
                            run_actions = r_action
                        min_wait_seconds = wait_seconds
        print 'min_wait_seconds: %s' % min_wait_seconds
        if min_wait_seconds > 0:
            time.sleep(min_wait_seconds)
            _now = datetime.now()
            print 'datetime: %s, run actions: %s' % (_now, run_actions)

def get_wait_seconds(fixtime, now):
    if fixtime < now:
        return 10000000000
    wait_seconds = (fixtime - now).total_seconds()
    return wait_seconds

if __name__ == '__main__':
    main()
