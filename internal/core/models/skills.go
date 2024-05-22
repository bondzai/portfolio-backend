package models

type Skill struct {
	ID        int    `json:"id"`
	ImageURL  string `json:"image_url"`
	IsShowing bool   `json:"is_showing"`
	Name      string `json:"name"`
	Topic     string `json:"topic"`
	URL       string `json:"url"`
}

var Skills = []Skill{
	{
		ID:        1,
		ImageURL:  "https://www.vectorlogo.zone/logos/golang/golang-icon.svg",
		IsShowing: true,
		Name:      "Go",
		Topic:     "language",
		URL:       "https://go.dev/",
	},
	{
		ID:        2,
		ImageURL:  "https://raw.githubusercontent.com/devicons/devicon/master/icons/javascript/javascript-original.svg",
		IsShowing: true,
		Name:      "JavaScript",
		Topic:     "language",
		URL:       "https://developer.mozilla.org/en-US/docs/Web/JavaScript/",
	},
	{
		ID:        3,
		ImageURL:  "https://raw.githubusercontent.com/devicons/devicon/master/icons/typescript/typescript-original.svg",
		IsShowing: false,
		Name:      "TypeScript",
		Topic:     "language",
		URL:       "https://www.typescriptlang.org/",
	},
	{
		ID:        4,
		ImageURL:  "https://www.vectorlogo.zone/logos/python/python-icon.svg",
		IsShowing: true,
		Name:      "Python",
		Topic:     "language",
		URL:       "https://www.python.org/",
	},
	{
		ID:        5,
		ImageURL:  "https://www.vectorlogo.zone/logos/gnu_bash/gnu_bash-icon.svg",
		IsShowing: true,
		Name:      "Shell",
		Topic:     "language",
		URL:       "https://www.gnu.org/software/bash/",
	},
	{
		ID:        6,
		ImageURL:  "https://res.cloudinary.com/dbdacfhye/image/upload/v1668240741/Portfolio/skills/PLC.svg",
		IsShowing: true,
		Name:      "PLC",
		Topic:     "language",
		URL:       "https://en.wikipedia.org/wiki/Programmable_logic_controller/",
	},
	{
		ID:        7,
		ImageURL:  "https://www.vectorlogo.zone/logos/reactjs/reactjs-icon.svg",
		IsShowing: true,
		Name:      "React",
		Topic:     "frontend",
		URL:       "https://reactjs.org/",
	},
	{
		ID:        8,
		ImageURL:  "https://www.gitbook.com/cdn-cgi/image/width=256,height=40,fit=contain,dpr=1,format=auto/https%3A%2F%2F373165937-files.gitbook.io%2F~%2Ffiles%2Fv0%2Fb%2Fgitbook-legacy-files%2Fo%2Fspaces%252F-M-XEvRn3rhe8BDVGkss%252Favatar-rectangle.png%3Fgeneration%3D1582298855816936%26alt%3Dmedia",
		IsShowing: true,
		Name:      "Fiber",
		Topic:     "backend",
		URL:       "https://docs.gofiber.io/",
	},
	{
		ID:        9,
		ImageURL:  "https://www.vectorlogo.zone/logos/nodejs/nodejs-icon.svg",
		IsShowing: true,
		Name:      "NodeJS",
		Topic:     "backend",
		URL:       "https://nodejs.org",
	},
	{
		ID:        10,
		ImageURL:  "https://raw.githubusercontent.com/devicons/devicon/master/icons/express/express-original-wordmark.svg",
		IsShowing: true,
		Name:      "Express",
		Topic:     "backend",
		URL:       "https://expressjs.com",
	},
	{
		ID:        11,
		ImageURL:  "https://www.vectorlogo.zone/logos/nestjs/nestjs-icon.svg",
		IsShowing: false,
		Name:      "Nest",
		Topic:     "backend",
		URL:       "https://nestjs.com/",
	},
	{
		ID:        12,
		ImageURL:  "https://www.vectorlogo.zone/logos/djangoproject/djangoproject-icon.svg",
		IsShowing: true,
		Name:      "Django",
		Topic:     "backend",
		URL:       "https://www.djangoproject.com/",
	},
	{
		ID:        13,
		ImageURL:  "https://cdn.worldvectorlogo.com/logos/fastapi.svg",
		IsShowing: false,
		Name:      "Fast",
		Topic:     "backend",
		URL:       "https://fastapi.tiangolo.com/",
	},
	{
		ID:        14,
		ImageURL:  "https://www.vectorlogo.zone/logos/graphql/graphql-icon.svg",
		IsShowing: true,
		Name:      "GraphQL",
		Topic:     "commu",
		URL:       "https://graphql.org/",
	},
	{
		ID:        15,
		ImageURL:  "https://www.vectorlogo.zone/logos/postgresql/postgresql-icon.svg",
		IsShowing: true,
		Name:      "PostreSQL",
		Topic:     "database",
		URL:       "https://www.postgresql.org/",
	},
	{
		ID:        16,
		ImageURL:  "https://www.vectorlogo.zone/logos/mysql/mysql-icon.svg",
		IsShowing: true,
		Name:      "MySQL",
		Topic:     "database",
		URL:       "https://www.mysql.com/",
	},
	{
		ID:        17,
		ImageURL:  "https://img.icons8.com/color/96/microsoft-sql-server.png",
		IsShowing: true,
		Name:      "SQLServer",
		Topic:     "database",
		URL:       "https://www.microsoft.com/en-us/sql-server/sql-server-downloads",
	},
	{
		ID:        18,
		ImageURL:  "https://www.vectorlogo.zone/logos/influxdata/influxdata-icon.svg",
		IsShowing: true,
		Name:      "Influx",
		Topic:     "database",
		URL:       "https://www.influxdata.com/",
	},
	{
		ID:        19,
		ImageURL:  "https://www.vectorlogo.zone/logos/mongodb/mongodb-icon.svg",
		IsShowing: true,
		Name:      "Mongo",
		Topic:     "database",
		URL:       "https://www.mongodb.com/",
	},
	{
		ID:        20,
		ImageURL:  "https://www.vectorlogo.zone/logos/firebase/firebase-icon.svg",
		IsShowing: false,
		Name:      "Firebase",
		Topic:     "database",
		URL:       "https://firebase.google.com/",
	},
	{
		ID:        22,
		ImageURL:  "https://www.vectorlogo.zone/logos/redis/redis-icon.svg",
		IsShowing: true,
		Name:      "Redis",
		Topic:     "database",
		URL:       "https://redis.io/",
	},
	{
		ID:        23,
		ImageURL:  "https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg",
		IsShowing: true,
		Name:      "Git",
		Topic:     "tools",
		URL:       "https://git-scm.com/",
	},
	{
		ID:        24,
		ImageURL:  "https://www.vectorlogo.zone/logos/github/github-icon.svg",
		IsShowing: true,
		Name:      "GitHub",
		Topic:     "tools",
		URL:       "https://github.com/",
	},
	{
		ID:        25,
		ImageURL:  "https://upload.wikimedia.org/wikipedia/commons/9/9a/Vmware.svg",
		IsShowing: true,
		Name:      "VMware",
		Topic:     "tools",
		URL:       "https://www.vmware.com/",
	},
	{
		ID:        26,
		ImageURL:  "https://www.vectorlogo.zone/logos/docker/docker-icon.svg",
		IsShowing: true,
		Name:      "Docker",
		Topic:     "tools",
		URL:       "https://www.docker.com/",
	},
	{
		ID:        28,
		ImageURL:  "https://www.vectorlogo.zone/logos/replit/replit-icon.svg",
		IsShowing: true,
		Name:      "Replit",
		Topic:     "tools",
		URL:       "https://replit.com/",
	},
	{
		ID:        30,
		ImageURL:  "https://www.vectorlogo.zone/logos/rabbitmq/rabbitmq-icon.svg",
		IsShowing: true,
		Name:      "RabbitMQ",
		Topic:     "commu",
		URL:       "https://www.rabbitmq.com/",
	},
	{
		ID:        33,
		ImageURL:  "https://www.vectorlogo.zone/logos/grafana/grafana-icon.svg",
		IsShowing: true,
		Name:      "Grafana",
		Topic:     "tools",
		URL:       "https://grafana.com/",
	},
	{
		ID:        35,
		ImageURL:  "https://www.vectorlogo.zone/logos/virtualbox/virtualbox-icon.svg",
		IsShowing: true,
		Name:      "Virtualbox",
		Topic:     "tools",
		URL:       "https://www.virtualbox.org/",
	},
	{
		ID:        37,
		ImageURL:  "https://www.vectorlogo.zone/logos/ubuntu/ubuntu-icon.svg",
		IsShowing: true,
		Name:      "Ubuntu",
		Topic:     "os",
		URL:       "https://ubuntu.com/",
	},
	{
		ID:        38,
		ImageURL:  "https://www.vectorlogo.zone/logos/archlinux/archlinux-icon.svg",
		IsShowing: true,
		Name:      "Arch",
		Topic:     "os",
		URL:       "https://archlinux.org/",
	},
	{
		ID:        39,
		ImageURL:  "https://www.vectorlogo.zone/logos/debian/debian-icon.svg",
		IsShowing: true,
		Name:      "Debian",
		Topic:     "os",
		URL:       "https://www.debian.org/",
	},
	{
		ID:        40,
		ImageURL:  "https://www.vectorlogo.zone/logos/raspberrypi/raspberrypi-icon.svg",
		IsShowing: true,
		Name:      "RaspberryPi",
		Topic:     "os",
		URL:       "https://www.raspberrypi.com/software/",
	},
	{
		ID:        42,
		ImageURL:  "https://www.vectorlogo.zone/logos/arduino/arduino-icon.svg",
		IsShowing: true,
		Name:      "Arduino",
		Topic:     "automation",
		URL:       "https://www.arduino.cc/",
	},
	{
		ID:        43,
		ImageURL:  "https://symbols-electrical.getvecta.com/stencil_262/71_rockwell-automation-icon.05b6277eb8.svg",
		IsShowing: true,
		Name:      "Rockwell",
		Topic:     "automation",
		URL:       "https://www.rockwellautomation.com/en-us/tools/software-subscriptions-updated.html",
	},
	{
		ID:        44,
		ImageURL:  "https://www.aveva.com/content/experience-fragments/aveva/en/site/header-2/master/_jcr_content/root/responsivegrid/globalheader/logo.coreimg.svg/1655394323761/header-logo.svg",
		IsShowing: true,
		Name:      "Aveva",
		Topic:     "automation",
		URL:       "https://www.aveva.com/en/solutions/operations/operations-control-hmi/",
	},
	{
		ID:        49,
		ImageURL:  "https://www.vectorlogo.zone/logos/grpcio/grpcio-ar21.svg",
		IsShowing: true,
		Name:      "gRPC",
		Topic:     "commu",
		URL:       "https://grpc.io/",
	},
}
