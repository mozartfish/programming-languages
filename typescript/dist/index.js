"use strict";
const authorOne = { name: 'mario', avatar: '/img/mario.png' };
const newPost = {
    title: 'my first post',
    body: 'this is content',
    tags: ['gaming', 'text'],
    create_at: new Date(),
    author: authorOne,
};
let age = 30;
let firstName = 'Mario';
let isFictional;
age = 31;
firstName = 'Luigi';
isFictional = false;
let planet = 'Earth';
let moons = 1;
let isLarge = false;
planet = 'Saturn';
moons = 145;
// null & unbdefined
let something;
let anotherThing;
// arrays
let names = [
    'harry potter',
    'lord voldemort',
    'dumbledore',
    'snape',
    'hagrid',
    'hermione',
    'ron',
];
let ages = [25, 24, 10];
names.push('this is a test');
ages.push(1);
//object literals
const g = names[2];
console.log(g);
//type inference with object literals
let user = {
    firstName: 'hello',
    age: 30,
    id: 10,
};
console.log(user);
