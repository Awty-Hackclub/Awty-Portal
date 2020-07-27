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
  - Tailwindcss
  - Vue Native (eventually)

### Backend

  - Python
  - Flask
  - Sqlalchemy (with some plain SQL)
  - MySQL
  - Redis
  - REST
  - UUID
  - Potential OAUTH if we implement 3rd party integration
  - Websocket/SocketIO (when we implement the chat)
  - Cookies (for auth; may rearchitect it though)
