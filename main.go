package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi"

	discord "github.com/bwmarrin/discordgo"
)

var messagesSent = make(map[string]time.Time)

var session *discord.Session

func main() {
	var err error
	session, err = discord.New("Bot " + ReadToken())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	err = session.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	RegisterCommands()

	err = session.UpdateStatus(0, "*aaaaaahhhh*")
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(ParseCommands)
	session.AddHandler(DidIHearSpook)

	StartSpookyWebServer()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-sc
}

func ReadToken() string {
	file := "token"
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(dat)
}

func DidIHearSpook(s *discord.Session, m *discord.MessageCreate) {
	if m.Author.Bot {
		return
	}
	content := m.Message.Content
	split := strings.Split(content, " ")
	didTheySpook := false
	for _, i := range split {
		if strings.ToLower(i) == "spook" || strings.ToLower(i) == "spooked" || strings.ToLower(i) == "spooks" {
			didTheySpook = true
		}
	}

	if didTheySpook {
		s.ChannelMessageSend(m.ChannelID, "Did I hear spook??")
	}
}

func StartSpookyWebServer() {
	r := chi.NewRouter()

	f, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Spook... could not compute...")
		return
	}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(f)
	})
	r.Get("/spook", func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("user")
		times, ok := messagesSent[user]
		if ok {
			if time.Since(times).Minutes() < 5 {
				http.Redirect(w, r, "/", 307)
				return
			}
		} else {
			messagesSent[user] = time.Now()
		}
		fmt.Println(user)
		if user == "" {
			fmt.Println("no user? :(")
			return
		}
		if session == nil {
			fmt.Println("wow i am spooked")
			return
		}
		id, err := session.UserChannelCreate(user)
		if err != nil || id == nil {
			return
		}
		session.ChannelMessageSend(id.ID, "Someone is thinking of you (from the internet)... And they decided to spook you. SPOOK!")
		http.Redirect(w, r, "/", 307)
	})
	http.ListenAndServe(":1337", r)
}
