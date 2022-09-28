# StatOn

Perhaps the easiest way to manage Status page with notifications for users.

## What it does

It was designed to be a status page for founders/creators, users, and nontech people also.

Founders - Easy to run and maintain, comes with a CRM for easy editing.
Users - users can subscribe in many ways as empowered by courier API.
Non-tech people - Cause the technical overhead is very less, anyone can work with it with ultra ease.
Besides all these, it's open source and Free

## Deployment

You'll need to run 3 things here, but before that run `npm install` in both strapi-backend and frontend.

In strapi-backend folder run

```bash
npm run develop
```

This will give you an UI to work with. Here i used demo@strapi.com as username and password as Courier@22

In frontend folder run

```bash
npm run dev
```

This will give you a front end page for getting info from strapi

In webhook-backend folder run

```bash
API_KEY="your_api_key" go run main.go
```

## Tech Stack

**Client:** Nuxt 3, TailwindCSS

**Server:** Strapi, Golang

## What's next for StatOps

We're still pretty lacking as I didn't get much time to work on it. Our immediate improvements on the chart are:

- Get Bulk mailing working
- Get appropriate colors for the Status page
- Provide a subscription option for user notifications (currently mailed only to me due to time constraints)
- Sorting Incidents component
- Show percentage of uptime
- Link incidents to Services
- Allow other subscription options like Discord, slack, etc.,
