package modules

import (
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/modules/bw"
	"github.com/kamushadenes/swayit/modules/chess"
	"github.com/kamushadenes/swayit/modules/crypto"
	externalIp "github.com/kamushadenes/swayit/modules/external_ip"
	"github.com/kamushadenes/swayit/modules/fan"
	"github.com/kamushadenes/swayit/modules/forex"
	"github.com/kamushadenes/swayit/modules/hash"
	"github.com/kamushadenes/swayit/modules/intel_gpu"
	"github.com/kamushadenes/swayit/modules/ip"
	"github.com/kamushadenes/swayit/modules/khal"
	"github.com/kamushadenes/swayit/modules/mail"
	"github.com/kamushadenes/swayit/modules/nordvpn"
	"github.com/kamushadenes/swayit/modules/notes"
	"github.com/kamushadenes/swayit/modules/opsgenie"
	"github.com/kamushadenes/swayit/modules/power"
	"github.com/kamushadenes/swayit/modules/prometheus"
	"github.com/kamushadenes/swayit/modules/snippets"
	"github.com/kamushadenes/swayit/modules/ssh"
	"github.com/kamushadenes/swayit/modules/timezone"
	"github.com/kamushadenes/swayit/modules/todoist"
	"github.com/kamushadenes/swayit/modules/uuidgen"
	"github.com/kamushadenes/swayit/modules/weather"
)

var Modules []common.Module

func RegisterModule(module common.Module) {
	Modules = append(Modules, module)
}

func init() {
	RegisterModule(chess.GetModule())
	RegisterModule(externalIp.GetModule())
	RegisterModule(fan.GetModule())
	RegisterModule(power.GetModule())
	RegisterModule(weather.GetModule())
	RegisterModule(mail.GetModule())
	RegisterModule(prometheus.GetModule())
	RegisterModule(crypto.GetModule())
	RegisterModule(forex.GetModule())
	RegisterModule(intel_gpu.GetModule())
	RegisterModule(khal.GetModule())
	RegisterModule(nordvpn.GetModule())
	RegisterModule(opsgenie.GetModule())
	RegisterModule(todoist.GetModule())
	RegisterModule(timezone.GetModule())
	RegisterModule(uuidgen.GetModule())
	RegisterModule(hash.GetModule())
	RegisterModule(ip.GetModule())
	RegisterModule(ssh.GetModule())
	RegisterModule(notes.GetModule())
	RegisterModule(snippets.GetModule())
	RegisterModule(bw.GetModule())
}
