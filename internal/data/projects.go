package data

type Project struct {
	HostURL     string `json:"host_url"`
	ID          int    `json:"id"`
	ImageURL    string `json:"image_url"`
	IsHighlight bool   `json:"is_highlight"`
	IsSleep     bool   `json:"is_sleep"`
	Language    string `json:"language"`
	Name        string `json:"name"`
	SourceURL   string `json:"source_url"`
	Status      string `json:"status"`
	Tools       string `json:"tools"`
}

var Projects = []Project{
	{
		HostURL:     "https://introbond-upload.cyclic.app/",
		ID:          1,
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
		ID:          2,
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
		ID:          3,
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
		ID:          4,
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
		ID:          5,
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
		ID:          6,
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
		ID:          7,
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
		ID:          8,
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
		ID:          9,
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
		ID:          10,
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
		ID:          11,
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
		ID:          13,
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
		ID:          14,
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