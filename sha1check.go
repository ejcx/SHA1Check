package main
import (
        "fmt"
        "log"
        "time"
        "flag"
        "crypto/tls"
        "crypto/x509"
)

const yellow= "Jan 1, 2016 at 0:01am (GMT)"
const red= "Jan 1, 2017 at 0:01am (GMT)"
func main() {
        host := flag.String("n", "", "-n example.com")
        port := flag.Int("p", -1, "-p 443")
        flag.Parse()
        if len(*host) == 0 {
                log.Fatalf("Hostname unset. Use -n")
        }
        if *port == -1 {
                *port = 443
        }
        connectto := fmt.Sprintf("%s:%d", *host, *port)
        conn, err := tls.Dial("tcp", connectto, &tls.Config{})
        if err != nil {
                log.Fatalf(err.Error())
        }
        yt, _ := time.Parse(yellow, "Jan 1, 2017 at 0:01am (GMT)")
        rt, _ := time.Parse(red, "Jan 1, 2017 at 0:01am (GMT)")
        st := conn.ConnectionState()
        chain := st.PeerCertificates
        chainyellow := false
        chainred := false
        for _, j := range chain {
                t := j.NotAfter
                yellow := t.After(yt)
                red := t.After(rt)
                names := ""
                for _, name := range j.DNSNames {
                        if names == "" {
                                names = name
                        } else {
                                fmt.Sprintf("%s, %s", names, name)
                        }
                }

                //If we don't have a DNSName, get the CommonName
                if names == "" {
                        names = j.Subject.CommonName
                }
                if (j.SignatureAlgorithm == x509.SHA1WithRSA){
                        if yellow {
                                fmt.Printf("\033[37m%s: \033[33mSecure with minor errors\n", names) 
                                chainyellow = chainyellow || yellow
                        } else if red {
                                fmt.Printf("\033[37m%s: \033[31mInsecure\n", names)
                                chainred = chainred || red
                        }
                } else {
                        fmt.Printf("\033[37m%s: \033[32mOk\n", names) 
                }
        }
        if chainred {
                fmt.Printf("\033[37m%s: \033[31mInsecure\n", *host)
        } else if chainyellow {
                fmt.Printf("\033[37m%s: \033[33mSecure with minor errors\n", *host) 
        } else {
                fmt.Printf("\033[32mSafe!\n")
        }
}
