package repositories

import (
	"github.com/bondzai/portfolio-backend/internal/models"
)

type (
	MockRepository interface {
		ReadCerts() ([]models.Certification, error)
		ReadProjects() ([]models.Project, error)
		ReadSkills() ([]models.Skill, error)
	}

	mockRepository struct{}
)

func NewMock() *mockRepository {
	return &mockRepository{}
}

func (mr *mockRepository) ReadCerts() ([]models.Certification, error) {
	certifications := []models.Certification{
		{
			Description: "Thin manager",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667635443/Portfolio/cirtifications/Cir-Thin.png",
			Name:        "Thin manager",
			OtherURL:    "",
		},
		{
			Description: "CSI Intouch",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667635441/Portfolio/cirtifications/Cir-CSI.png",
			Name:        "CSI Intouch",
			OtherURL:    "",
		},
		{
			Description: "Pre-CCNA",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667635442/Portfolio/cirtifications/Cir-Network.jpg",
			Name:        "Pre-CCNA",
			OtherURL:    "",
		},
		{
			Description: "DevLab3",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667635442/Portfolio/cirtifications/Cir-DevLab.png",
			Name:        "DevLab3",
			OtherURL:    "",
		},
		{
			Description: "Backend",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667635442/Portfolio/cirtifications/Cir-Udemy-Backend.jpg",
			Name:        "Backend",
			OtherURL:    "",
		},
		{
			Description: "Protocols",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667635442/Portfolio/cirtifications/Cir-Udemy-Protocols.jpg",
			Name:        "Protocols",
			OtherURL:    "",
		},
		{
			Description: "Go",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667635442/Portfolio/cirtifications/Cir-Udemy-Backend-Master.jpg",
			Name:        "Go",
			OtherURL:    "",
		},
		{
			Description: "ChatGPT",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1689411286/Portfolio/cirtifications/Cir-ChatGPT.png",
			Name:        "ChatGPT",
			OtherURL:    "",
		},
		{
			Description: "Apollo GraphQL",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1694930735/Portfolio/cirtifications/Cir-Apollo-Associate.png",
			Name:        "Apoll-Associate",
			OtherURL:    "https://www.apollographql.com/tutorials/certifications/03d75505-810a-4c6a-a9d1-85d18a27b941",
		},
	}

	return certifications, nil
}

func (mr *mockRepository) ReadProjects() ([]models.Project, error) {
	projects := []models.Project{
		{
			HostURL:     "https://introbond-upload.cyclic.app/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634064/Portfolio/project-cloudinary.png",
			IsHighlight: false,
			IsSleep:     false,
			Language:    "Javascript",
			Name:        "Microservice: File Upload",
			SourceURL:   "https://github.com/introbond/Lab-express-cloudinary",
			Status:      "offline",
			Tools:       "",
		},
		{
			HostURL:     "https://bondzai.github.io/micro-app-quiz/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634064/Portfolio/project-quiz.png",
			IsHighlight: false,
			IsSleep:     false,
			Language:    "Javascript",
			Name:        "Quiz",
			SourceURL:   "https://github.com/bondzai/micro-app-quiz/tree/main",
			Status:      "online",
			Tools:       "",
		},
		{
			HostURL:     "https://bondzai.github.io/micro-app-todolist/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634065/Portfolio/project-todolist.png",
			IsHighlight: false,
			IsSleep:     false,
			Language:    "Javascript",
			Name:        "To-Do List",
			SourceURL:   "https://github.com/bondzai/micro-app-todolist/tree/main",
			Status:      "online",
			Tools:       "",
		},
		{
			HostURL:     "https://bondzai.github.io/micro-app-contries-search-filter/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634064/Portfolio/project-contries-info.png",
			IsHighlight: false,
			IsSleep:     false,
			Language:    "Javascript",
			Name:        "Contries Information",
			SourceURL:   "https://github.com/bondzai/micro-app-contries-search-filter/tree/main",
			Status:      "online",
			Tools:       "",
		},
		{
			HostURL:     "https://microservice-timestamp.introbond.repl.co",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634065/Portfolio/project-timestamp.png",
			IsHighlight: false,
			IsSleep:     true,
			Language:    "Javascript",
			Name:        "Microservice: Time Stamp",
			SourceURL:   "https://replit.com/@introbond/microservice-timestamp#index.js",
			Status:      "offline",
			Tools:       "",
		},
		{
			HostURL:     "https://bondzai.github.io/micro-app-simple-calculator/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634064/Portfolio/project-calculator.png",
			IsHighlight: false,
			IsSleep:     false,
			Language:    "Javascript",
			Name:        "Calculator",
			SourceURL:   "https://github.com/bondzai/micro-app-api-currency-exchange/tree/main",
			Status:      "online",
			Tools:       "",
		},
		{
			HostURL:     "https://microservice-headerparser.introbond.repl.co/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634064/Portfolio/project-req-parser.png",
			IsHighlight: false,
			IsSleep:     true,
			Language:    "Javascript",
			Name:        "Microservice: Request Parser",
			SourceURL:   "https://replit.com/@introbond/microservice-headerparser#index.js",
			Status:      "offline",
			Tools:       "",
		},
		{
			HostURL:     "https://introbond-crypto-ui.onrender.com/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634064/Portfolio/project-hash.png",
			IsHighlight: false,
			IsSleep:     true,
			Language:    "Javascript",
			Name:        "Microservice: SHA-256",
			SourceURL:   "https://github.com/bondzai/micro-app-crypto-ui",
			Status:      "online",
			Tools:       "",
		},
		{
			HostURL:     "https://bondzai.github.io/micro-app-api-currency-exchange/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634065/Portfolio/project-exchange.png",
			IsHighlight: false,
			IsSleep:     false,
			Language:    "Javascript",
			Name:        "Currency Exchange",
			SourceURL:   "https://github.com/bondzai/micro-app-api-currency-exchange/tree/main",
			Status:      "online",
			Tools:       "",
		},
		{
			HostURL:     "https://ecommerce.introbond.repl.co/api-docs/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1667634064/Portfolio/project-ecommerce.png",
			IsHighlight: false,
			IsSleep:     true,
			Language:    "Javascript",
			Name:        "API: E-Commerce",
			SourceURL:   "https://replit.com/@introbond/eCommerce?v=1",
			Status:      "offline",
			Tools:       "",
		},
		{
			HostURL:     "https://swiftdev.onrender.com/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1681039500/project-scrum.png",
			IsHighlight: false,
			IsSleep:     false,
			Language:    "Javascript",
			Name:        "Scrum Dashboard",
			SourceURL:   "https://github.com/bondzai/scrum-dashboard",
			Status:      "offline",
			Tools:       "",
		},
		{
			HostURL:     "https://apollo-catstronauts-client.vercel.app/",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1694937857/Portfolio/projects/project-cats-client.png",
			IsHighlight: true,
			IsSleep:     true,
			Language:    "Javascript",
			Name:        "GraphQL-Catstronauts-client",
			SourceURL:   "https://github.com/bondzai/apollo-catstronauts-client",
			Status:      "online",
			Tools:       "",
		},
		{
			HostURL:     "https://apollo-catstronauts-server.introbond.repl.co",
			ImageURL:    "https://res.cloudinary.com/dbdacfhye/image/upload/v1694937857/Portfolio/projects/project-cats-server.png",
			IsHighlight: true,
			IsSleep:     true,
			Language:    "Javascript",
			Name:        "GraphQL-Catstronauts-server",
			SourceURL:   "https://github.com/bondzai/apollo-catstronauts-server",
			Status:      "offline",
			Tools:       "",
		},
	}

	return projects, nil
}

func (mr *mockRepository) ReadSkills() ([]models.Skill, error) {
	skills := []models.Skill{
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/golang/golang-icon.svg",
			IsShowing: true,
			Name:      "Go",
			Topic:     "language",
			URL:       "https://go.dev/",
		},
		{
			ImageURL:  "https://raw.githubusercontent.com/devicons/devicon/master/icons/javascript/javascript-original.svg",
			IsShowing: true,
			Name:      "JavaScript",
			Topic:     "language",
			URL:       "https://developer.mozilla.org/en-US/docs/Web/JavaScript/",
		},
		{
			ImageURL:  "https://raw.githubusercontent.com/devicons/devicon/master/icons/typescript/typescript-original.svg",
			IsShowing: false,
			Name:      "TypeScript",
			Topic:     "language",
			URL:       "https://www.typescriptlang.org/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/python/python-icon.svg",
			IsShowing: true,
			Name:      "Python",
			Topic:     "language",
			URL:       "https://www.python.org/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/gnu_bash/gnu_bash-icon.svg",
			IsShowing: true,
			Name:      "Shell",
			Topic:     "language",
			URL:       "https://www.gnu.org/software/bash/",
		},
		{
			ImageURL:  "https://res.cloudinary.com/dbdacfhye/image/upload/v1668240741/Portfolio/skills/PLC.svg",
			IsShowing: true,
			Name:      "PLC",
			Topic:     "language",
			URL:       "https://en.wikipedia.org/wiki/Programmable_logic_controller/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/reactjs/reactjs-icon.svg",
			IsShowing: true,
			Name:      "React",
			Topic:     "frontend",
			URL:       "https://reactjs.org/",
		},
		{
			ImageURL:  "https://www.gitbook.com/cdn-cgi/image/width=256,height=40,fit=contain,dpr=1,format=auto/https%3A%2F%2F373165937-files.gitbook.io%2F~%2Ffiles%2Fv0%2Fb%2Fgitbook-legacy-files%2Fo%2Fspaces%252F-M-XEvRn3rhe8BDVGkss%252Favatar-rectangle.png%3Fgeneration%3D1582298855816936%26alt%3Dmedia",
			IsShowing: true,
			Name:      "Fiber",
			Topic:     "backend",
			URL:       "https://docs.gofiber.io/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/nodejs/nodejs-icon.svg",
			IsShowing: true,
			Name:      "NodeJS",
			Topic:     "backend",
			URL:       "https://nodejs.org",
		},
		{
			ImageURL:  "https://raw.githubusercontent.com/devicons/devicon/master/icons/express/express-original-wordmark.svg",
			IsShowing: true,
			Name:      "Express",
			Topic:     "backend",
			URL:       "https://expressjs.com",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/nestjs/nestjs-icon.svg",
			IsShowing: false,
			Name:      "Nest",
			Topic:     "backend",
			URL:       "https://nestjs.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/djangoproject/djangoproject-icon.svg",
			IsShowing: true,
			Name:      "Django",
			Topic:     "backend",
			URL:       "https://www.djangoproject.com/",
		},
		{
			ImageURL:  "https://cdn.worldvectorlogo.com/logos/fastapi.svg",
			IsShowing: false,
			Name:      "Fast",
			Topic:     "backend",
			URL:       "https://fastapi.tiangolo.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/graphql/graphql-icon.svg",
			IsShowing: true,
			Name:      "GraphQL",
			Topic:     "commu",
			URL:       "https://graphql.org/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/postgresql/postgresql-icon.svg",
			IsShowing: true,
			Name:      "PostreSQL",
			Topic:     "database",
			URL:       "https://www.postgresql.org/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/mysql/mysql-icon.svg",
			IsShowing: true,
			Name:      "MySQL",
			Topic:     "database",
			URL:       "https://www.mysql.com/",
		},
		{
			ImageURL:  "https://img.icons8.com/color/96/microsoft-sql-server.png",
			IsShowing: true,
			Name:      "SQLServer",
			Topic:     "database",
			URL:       "https://www.microsoft.com/en-us/sql-server/sql-server-downloads",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/influxdata/influxdata-icon.svg",
			IsShowing: true,
			Name:      "Influx",
			Topic:     "database",
			URL:       "https://www.influxdata.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/mongodb/mongodb-icon.svg",
			IsShowing: true,
			Name:      "Mongo",
			Topic:     "database",
			URL:       "https://www.mongodb.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/firebase/firebase-icon.svg",
			IsShowing: false,
			Name:      "Firebase",
			Topic:     "database",
			URL:       "https://firebase.google.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/redis/redis-icon.svg",
			IsShowing: true,
			Name:      "Redis",
			Topic:     "database",
			URL:       "https://redis.io/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg",
			IsShowing: true,
			Name:      "Git",
			Topic:     "tools",
			URL:       "https://git-scm.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/github/github-icon.svg",
			IsShowing: true,
			Name:      "GitHub",
			Topic:     "tools",
			URL:       "https://github.com/",
		},
		{
			ImageURL:  "https://upload.wikimedia.org/wikipedia/commons/9/9a/Vmware.svg",
			IsShowing: true,
			Name:      "VMware",
			Topic:     "tools",
			URL:       "https://www.vmware.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/docker/docker-icon.svg",
			IsShowing: true,
			Name:      "Docker",
			Topic:     "tools",
			URL:       "https://www.docker.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/replit/replit-icon.svg",
			IsShowing: true,
			Name:      "Replit",
			Topic:     "tools",
			URL:       "https://replit.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/rabbitmq/rabbitmq-icon.svg",
			IsShowing: true,
			Name:      "RabbitMQ",
			Topic:     "commu",
			URL:       "https://www.rabbitmq.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/grafana/grafana-icon.svg",
			IsShowing: true,
			Name:      "Grafana",
			Topic:     "tools",
			URL:       "https://grafana.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/virtualbox/virtualbox-icon.svg",
			IsShowing: true,
			Name:      "Virtualbox",
			Topic:     "tools",
			URL:       "https://www.virtualbox.org/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/ubuntu/ubuntu-icon.svg",
			IsShowing: true,
			Name:      "Ubuntu",
			Topic:     "os",
			URL:       "https://ubuntu.com/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/archlinux/archlinux-icon.svg",
			IsShowing: true,
			Name:      "Arch",
			Topic:     "os",
			URL:       "https://archlinux.org/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/debian/debian-icon.svg",
			IsShowing: true,
			Name:      "Debian",
			Topic:     "os",
			URL:       "https://www.debian.org/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/raspberrypi/raspberrypi-icon.svg",
			IsShowing: true,
			Name:      "RaspberryPi",
			Topic:     "os",
			URL:       "https://www.raspberrypi.com/software/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/arduino/arduino-icon.svg",
			IsShowing: true,
			Name:      "Arduino",
			Topic:     "automation",
			URL:       "https://www.arduino.cc/",
		},
		{
			ImageURL:  "https://symbols-electrical.getvecta.com/stencil_262/71_rockwell-automation-icon.05b6277eb8.svg",
			IsShowing: true,
			Name:      "Rockwell",
			Topic:     "automation",
			URL:       "https://www.rockwellautomation.com/en-us/tools/software-subscriptions-updated.html",
		},
		{
			ImageURL:  "https://www.aveva.com/content/experience-fragments/aveva/en/site/header-2/master/_jcr_content/root/responsivegrid/globalheader/logo.coreimg.svg/1655394323761/header-logo.svg",
			IsShowing: true,
			Name:      "Aveva",
			Topic:     "automation",
			URL:       "https://www.aveva.com/en/solutions/operations/operations-control-hmi/",
		},
		{
			ImageURL:  "https://www.vectorlogo.zone/logos/grpcio/grpcio-ar21.svg",
			IsShowing: true,
			Name:      "gRPC",
			Topic:     "commu",
			URL:       "https://grpc.io/",
		},
	}

	return skills, nil
}
