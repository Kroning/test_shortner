# test_shortner
Contacts: https://t.me/Kroning https://www.linkedin.com/in/alexander-bulakhov-3122aa62/ <br>

The best way to learn is experimenting.<br>
I am trying to build shortner (or something alike) with 2 services (admin, redirect).<br>
It's not a good idea, but for learning purpose I decided to make 2 services at 1 repo: webinterface + redirector. In production it's better to place services in different modules to deploy easily on different servers.

If you can point something that I am doing wrong - feel free to write.

TODO:
1. HTML is simple. Better way - some API calls with JSON in return.
2. Logging
3. Error handling
4. Database (https://github.com/jackc/pgx)
5. Routing module (because I can)

If there any better way to pass db handler to httpHandler? (Except main -> App struct -> Page struct)

Tried already:<br>
Modules & packages<br>
Project Layout<br>
net/http, basic routing, handlers<br>
Struct embedding<br>
Html/template basics<br>
