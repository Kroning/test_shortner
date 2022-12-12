# test_shortner
Contacts: https://t.me/Kroning https://www.linkedin.com/in/alexander-bulakhov-3122aa62/ <br>

Experimenting with libs and technologies.<br>
I am trying to build shortner (or something alike) with 2 services (admin, redirect). Not the best way, but best for learning new :)<br>

If you can point something that I am doing wrong - feel free to write.

TODO:
1. HTML is simple. Better way - some API calls with JSON in return.
2. Routing module (because I can).
3. How to install (sql, nginx, rc.d)
4. Authorization and private links.
5. Link's statistic.

If there any better way to pass db handler to httpHandler? (Except main -> App struct -> Page struct)

Tried already:<br>
Modules & packages<br>
Project Layout<br>
net/http, basic routing, handlers<br>
Struct embedding<br>
Html/template basics<br>
Docs (commented, read with m=all)<br>
Configs (yml & env)<br>
Logging (log package, saving to file)<br>
Error handling (basic, added where it needs to return error)<br>
Tests (just a little, not 100% coverage; testify, require, assert)
Database (pgxpool)
Flags

v.1.0.* : https://github.com/Kroning/test_shortner/blob/master/install/intall-1.0.md<br>

