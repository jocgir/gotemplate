Function ls2
No param   : @color("green", ls2())
With param : @color("yellow", ls2("/"))

Function xxx
@xxx := 123456
No param   : @color("green", xxx("Martine", "Bourque"))
With param : @color("yellow", xxx("Jocelyn Giroux"))

Function myfunc
no param   : @color("green", myfunc())
With param : @color("yellow", myfunc("/"))

Call run
@run("ls")
