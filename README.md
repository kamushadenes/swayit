<h1 align="center">SwayIT</h1>
<p align="center">
A set of modules to provide at-glance information, focused on Sway and WayBar
</p>

<hr> 

## Install

```sh
go get github.com/kamushadenes/swayit
```

## Usage

```sh
swayit help
```

List available modules
```sh
swayit list
```

Run a module to make it's output available
```sh
swayit run <module>
```

Get a module output (WayBar formatted JSON)
```sh
swayit output <module>
```

Run a module command
```sh
swayit command <module> [args]
```

## Modules

| Name       | Slug       | Description                    |
| ---------- | ---------- | ------------------------------ |
| BitWarden  | bw         | BitWarden Credential Picker    |
| Chess      | chess      | Chess.com Notifications        |
| Crypto     | crypto     | CryptoCoin Ticker              |
| ExternalIP | externalIp | Get External IP                |
| Fan        | fan        | Laptop Fan RPM                 |
| Forex      | forex      | FIAT Ticker                    |
| Hash       | hash       | Quick Hash Menu                |
| IP         | ip         | IP Information Menu            |
| IntelGPU   | intelGpu   | Intel GPU Top                  |
| Khal       | khal       | Khal Event Integration         |
| Mail       | mail       | Mail Notifications             |
| NordVPN    | nordvpn    | NordVPN Status                 |
| Notes      | notes      | Note Taking                    |
| OpsGenie   | opsgenie   | OpsGenie On-Call Notifications |
| Power      | power      | Power Monitor                  |
| Prometheus | prometheus | Prometheus Alarms              |
| SSH        | ssh        | SSH Menu                       |
| Snippets   | snippets   | Snippets Menu                  |
| Timezone   | timezone   | Timezone Information           |
| Todoist    | todoist    | Todoist Tasks                  |
| UUIDGen    | uuidgen    | UUID Generation Menu           |
| Weather    | weather    | Weather Notifications          |

## Author

👤 **Kamus Hadenes**

* Website: https://blog.hadenes.io
* Twitter: [@kamushadenes](https://twitter.com/kamushadenes)
* Github: [@kamushadenes](https://github.com/kamushadenes)
* LinkedIn: [@kamushadenes](https://linkedin.com/in/kamushadenes)
