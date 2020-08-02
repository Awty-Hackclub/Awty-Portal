# Awty-Portal

A social media app inspired by Twitter.

## Functionality

### MVP functionality

- User accounts
- User authentication (with JWT)
- Ability to post
- Ability to comment on posts
- Post likes
- Notification system
- Following users

### Eventual functionality

- Direct messages
- Media uploading (images and videos)
- Reposting (equivalent to retweets)
- Comment likes
- Mobile (via vue native) and desktop version (via electronjs)
- tagging functionality (@ and hashtags)
- Recommendation system powered by ML
- "Trending" recommendations
- Livestreaming? There is potential for this but it will require a lot of architectural accomodation.
- Mobile port
- Editing posts/comments
- Changing site theme?

## Tech Stack

### DevOps

- Docker
- Docker-Compose
- Bash
- Potential Docker-Swarm if we deploy

### Frontend

- VueJS
- VueX
- (possible nuxtjs)
- Vue Native (eventually)
- Figma (for UI/UX prototypes)

### Backend

- Python
- Flask
- SqlAlchemy (with some plain SQL)
- MySQL
- JWT auth
- REST
- UUID
- Potential OAUTH if we implement 3rd party integration
- Websocket/SocketIO (when we implement the chat)
- Cookies (for auth; may rearchitect it though)
- RabbitMQ (potential)

## How to run the app

- `./run.sh` *for linux users*
- *windows TBD* (most people implementing this app on windows will be on WSL so it shouldn't be a problem)
