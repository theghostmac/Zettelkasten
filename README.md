# Zettelkasten

Zettelkasten is a powerful storage site designed for developers to create accounts, log in, and organize their thoughts, 
slips, and ideas in the form of notes called "zettels." This application is built with Golang, a fast and 
efficient programming language, to ensure optimal performance and scalability.

## Features

- **Account Creation**: Users can easily create new accounts by providing their desired username and password.
- **Login**: Registered users can securely log in to their accounts using their credentials.
- **Zettel Creation**: Once logged in, users can write and save zettels containing their thoughts, slips, or ideas. Each zettel can be given a title and a body.
- **Organization and Tagging**: Zettels can be organized using tags, allowing users to categorize and group related notes together for easier navigation and retrieval.
- **Search**: The app provides a powerful search functionality, enabling users to find specific zettels based on titles, tags, or content.
- **Editing and Deleting**: Users have the flexibility to edit or delete their zettels whenever necessary, ensuring they can refine their ideas over time.
- **User Profiles**: Each user has a dedicated profile page, displaying their username and a collection of their zettels.
- **Collaboration**: Users can choose to share their zettels with others, enabling collaboration and idea sharing within the developer community.

## Installation

To install and run Zettelkasten locally, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/zettelkasten.git`
2. Navigate to the project directory: `cd zettelkasten`
3. Install the required dependencies: `go get -d ./...`
4. Build the application: `go build`
5. Run the application: `./zettelkasten`

By default, the application will be accessible at `http://localhost:8082`, but you can use a port of your choice with:
```shell
go run ./cmd/http -addr=":<choice>"
```

## Configuration (Coming soon 🔜)

Zettelkasten can be customized by modifying the configuration file located at `config.yaml`. This file allows you to specify 
settings such as database credentials, server port, and other application-specific options.

## Contributing

We welcome contributions from the community to enhance Zettelkasten and make it even better. To contribute, please follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch for your feature or bug fix.
3. Make the necessary changes and additions.
4. Commit and push your changes to your forked repository.
5. Submit a pull request to the main repository, explaining your changes and why they should be merged.

Please ensure that your code adheres to our coding standards and is accompanied by appropriate tests.

## License

Zettelkasten is released under the [MIT License](LICENSE). You are free to use, modify, and distribute this software as per the terms of the license.

## Acknowledgements

We would like to thank the open-source community for their invaluable contributions to the tools and libraries used in building Zettelkasten.

## Contact

If you have any questions, suggestions, or feedback, please feel free to reach out to us at [macbobbychibuzor@gmail.com](mailto:macbobbychibuzor@gmail.com). 
We appreciate your interest in Zettelkasten and look forward to hearing from you!