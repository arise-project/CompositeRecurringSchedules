# Composite Recurring Schedules
we will need quite the complex/powerful scheduling architecture: billing/invoicing, social media posts, appointments/meetings, alerts, newsletters, tasks/reminders... all need a very flexible recurring capacity... 


almost like there should be a Scheduling Scripting Language... {one time/every} {#} 
{seconds/minutes/hours/days/weeks/months/years/date/weekday/weekend/days of week} {starting date/time} {action} ... would be good if you could save that as a 'schedule' object and then make rules that stack.. like

{every} {1} {dec 25} {starting dec 25 2016} {do nothing} = $christmas
{every} {2} {weeks} {starting jan 1 2016} {send email to list} EXCEPT {$Christmas} = biweeklynewsletter

or better yet
{every} {1} {jan 1} {starting jan 1 2016} {do nothing} = $newyears
$holidays=$Christmas+$newyears

{every} {1} {week} {starting jan 1 2016} {email list} EXCEPT {$holidays} 

etc

http://martinfowler.com/apsupp/recurring.pdf


I'd want to solve all the going further part of the event paper. I'd prefer to stay away from java and things we can't easily modify

Beyond a solid structure id want a nice interface and a scripting ability and composition of "calendars"/schedule sets. 
It would be the most flexible way to handle any kind of odd event schedule and making it just emit events that other services can subscribe 
to will make it useful to handle everything time related in all of our services

Github... Id like to see some kind of algorithm and object diagram and swagger rest definition. 
Perhaps class relation and sequence and logic flowchart or maybe just a solid description of your solution ...

