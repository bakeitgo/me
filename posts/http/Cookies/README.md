### Cookies

* Small piece of data which server sends to the browser, and browser stores it.

* Browser sends all cookies with all next requests to the server which originated cookie.

* Cookie is domain specific.

#### How to set a cookie?

* Cookie can be set either client-side and server-side.

* Server-side the Set-Cookie=value is sent along with response due to request, for e.g. sign in

* Client-side it can be done with JS, for e.g. -- document.cookie="aaa=value

#### How to remove a cookie?

* Cookie can be removed either client-side and server-side

* Server-side:
    * Via specifying Set-Cookie attribute Expires=Date || Max-Age=timestamp
    * The further is less error prone due to easier parse, and takes precedence on Expires if both used.

* Session cookie
    * Cookie is considered as "Session" one when neither Expires / Max-Age is not set.

#### How to update a cookie?

* Cookie can be updated either client-side and server-side

* Server-side, via sending Set-Cookie: CookieToUpdate=NewValue header.

* Client-side, via JS -- document.cookie="CookieToUpdate="NewValue"
    * Consider that cookie cant be updated if HttpOnly attribute has been set to it. 
    * You cant modify cookie directly in JS fetch / XMLHttpRequest API

#### Security considerations for cookies:

* Attribute: Secure -> 
    * Client -> Server -- cookie will be sent only if connection is on top of TLS (https protocol),
    * Server -> Client -- cookie will be not set if its in insecure connection

* Attribute: HttpOnly -> 
    * With this attr, cookie cant be modified with JS, only after reaching the server.

* Attribute: Domain -> Specifies which server can receive cookie (based on domain)
    * By default (if its not set), Cookie can only be received by the cookie issuer (the domain from cookie originates from)
    * Setting Domain, makes cookie accessible from subdomains.

* Attribute: Path -> Specifies on which paths cookie will be sent.
    * It match subpaths aswell so "Path=/", will also match /user, /user/cart etc.

* Attribute: SameSite -> Restricts cookies usage across Cross-origin requests (request source is different than target source.
    * SameSite=Strict -> Only when source origin matches dest origin, the cookie will be sent.
    * SameSite=Lax -> Only when source origin matches dest origin AND the cross-origin will send request to cookie issuer domain, the cookie will be sent.
    * SameSite=None -> Cookie will be sent when source origin matches dest, cross-origin will send request to our cookie issuer domain, AND when cross-origin request resources from cookie issuer domain (To make it works on this prop also the Secure attribute must be set)


#### Cookie prefixes

* By default, cookie can be changed, for e.g. bad actor will add Domain attribute, which will lead to cookie being used on subdomains, and therefore can make for e.g. session fixatio attack. There is why cookie prefixes has come from.

* \_\_Host-: with this prefix, the cookie is accepted to being created in Set-Cookie only if it has attributes:
    * **Secure**, DOES NOT INCLUDE **Domain**, and **Path** is set to */*
    * ***THIS COOKIE IS DOMAIN-LOCKED***

* \_\_Secure-: it is accepted in Set-Cookie only if it has **Secure** attribute

* From now subdomain cookies cannot affect whole domain

#### Cookie-related regulations

* Due to GDPR / ePrivacy Directive / California Consumer Privacy Act, we need to inform user about cookies we store, and make them decide to agree or disagree for storing them



