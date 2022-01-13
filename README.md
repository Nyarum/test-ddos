Task:
Design and implement “Word of Wisdom” tcp server.   
• TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.   
• The choice of the POW algorithm should be explained.   
• After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.   
• Docker file should be provided both for the server and for the client that solves the POW challenge

Also would be great to have:
- Client test 
- Describe of choosed procotol
- Separated client package
- Comments of code

Additionals from author:
- References were used: wikipedia, github.com/bwesterb/go-pow, in result I choose library written on golang because it saves time to write implementation and it looking very friendly to use it
- Procotol based on std bufio package and ping-pong exchange way
- Used clean architecture following go-way style
- Quotes I took from https://everydaypower.com/words-of-wisdom/

How to run:
- make docker && make docker/create/network (Building images and creates local network)
- make docker/run/server (Running tcp server)
- make docker/run/client (Running tcp client and connects to tcp server)