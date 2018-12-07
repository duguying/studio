package p2p

import (
	"context"
	"crypto/rand"
	"duguying/studio/g"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-net"
	"github.com/multiformats/go-multiaddr"
	"log"
)

func Init() {
	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	if err != nil {
		log.Println("generate key failed, err:", err.Error())
		return
	}

	// generate p2p address
	addr := fmt.Sprintf("/ip4/%s/tcp/%d", g.Config.Get("p2p", "host", "0.0.0.0"), g.Config.GetInt64("p2p", "port", 23001))
	sourceMultiAddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		log.Println("generate address failed, err:", err.Error())
		return
	}

	// store p2p address as global
	g.P2pAddr = addr

	// create p2p host
	ctx := context.Background()
	host,err:=libp2p.New(
		ctx,
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)
	if err != nil {
		log.Println("create p2p host failed, err:", err.Error())
		return
	}

	host.SetStreamHandler("/shell", func(stream net.Stream) {

	})
}
