Entity inside the Domain Layer:
1. We separate the Entity from its representation in the database.
2. Its purpose to encapsulate essential business logic models.
3. Entities do not necessarily replicate the database structure.
4. Here we validate entities

Storage inside the Domain Layer:
Seems like it must NOT be here. 
Because we can put Storage interface to the service /
If we have decorators in Storage we additionally put interface there.

But here is the problem for developers, it stops parallel development and reduce feature delivery.
Our plan is to speed the proccess up.
We have couple tasks for developers which in result lead us to the fully working service.
The first developer creates domain part of the service with all entities and storage interface. Then others are able to develop: service, storage and etc without blockers.