# test_shortner

The best way to learn is experimenting.
I am trying to build shortner (or something alike) with 2 services (admin, redirect).
It's not a good idea, but for learning purpose I decided to make 2 services at 1 repo: webinterface + redirector. In production it's better to place services in different modules to deploy easily on different servers.

If you can point something that I am doing wrong - feel free to write.

TODO:
1. HTML is simple. Better way - some API calls with JSON in return.
2. Logging
3. Error handling

If there any better way to pass db handler to httpHandler? (Except main -> App struct -> Page struct)
