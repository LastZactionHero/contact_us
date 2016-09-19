# Simple "Contact Us" Form

## Environmental Variables

- CONTACT_US_DB_USER - Mysql DB User
- CONTACT_US_DB_PASS - Mysql DB Password
- CONTACT_US_DB_NAME - Mysql DB Name
- CONTACT_US_PORT - Server Port

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
