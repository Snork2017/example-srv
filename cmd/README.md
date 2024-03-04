Handlers belong to cmd/app:
1. Each app has its own handlers and configs. 
2. We pawn the architecture for the future in case if we need more apps in cmd with its own handlers and configs.
3. If somebody comes to see the repo then he will see which handler belongs to which app. We do NOT confuse other developers keeping all handlers in internal, they dont need to waste their time.
4. We keep only usable configs for apps.
