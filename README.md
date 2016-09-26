# Simple "Contact Us" Form

## Environmental Variables

- CONTACT_US_DB_USER - Mysql DB User
- CONTACT_US_DB_PASS - Mysql DB Password
- CONTACT_US_DB_NAME - Mysql DB Name
- CONTACT_US_PORT - Server Port
- CONTACT_US_IFTTT_KEY - IFTTT Maker Key
- CONTACT_US_IFTTT_TRIGGER - IFTTT Trigger Name

## Create Contact
POST: /contact

Body:
```
{
  email: 'user@site.com', // string
  name: 'Jeff Userson',   // string
  phone: '333-111-2234',  // string
  message: 'Hello world!' // string
}
```

Response: none

# Create Skill
POST: /skills

Body:
```
{
  name: 'C++' // string
}
```

Response: none (204)

# List Skills
GET: /skills

Response: (200)
```
[
  {
    name: 'C++',
  },
  {
    name: 'F#'
  }
]
```

# Create Contractor
POST: /contractor

Body:
```
{
  name: 'Jeff Contractor', // string,
  city: 'Boulder', // string,
  phone: '303-555-1123', // string
  currently_employed: true, // boolean
  availability: 'full_time', // string
  skills: [1,2,3], // array of Skill IDs,
  projects: 'all sorts of stuff...', // string
  github: 'http://www.github.com...', // string
  linkedin: 'http://www.linkedin.com...', // string
  website: 'http://www.personaliste.com', // string
  anything_else: 'I am awesome', // string
}
```

Response: none (204)
