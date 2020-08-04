# Awty-Portal

A social media app inspired by Twitter.

## Functionality

### Currently implemented functionality

- None

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
- Docker-Swarm if we deploy (potential)

### Frontend

- ReactJS 
- Redux (potential)
- React Native (eventually/potential)
- Figma (for UI/UX prototypes)

### Backend

- Python
- Flask
- SqlAlchemy (potential); may just use hardcoded queries
- MySQL
- JWT auth
- REST
- UUID
- Websocket/SocketIO (when we implement the chat)
- Cookies (for auth; may rearchitect it though)
- RabbitMQ (eventual for notifications, webhooks, etc)

## How to run the app

- `./run.sh` *for linux users*
- *windows TBD* (most people implementing this app on windows will be on WSL so it shouldn't be a problem)
