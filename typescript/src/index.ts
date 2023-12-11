// interfaces
interface Author {
  name: string;
  avatar: string;
}
const authorOne: Author = { name: 'mario', avatar: '/img/mario.png' };
interface Post {
  title: string;
  body: string;
  tags: string[];
  create_at: Date;
  author: Author;
}

const newPost: Post = {
  title: 'my first post',
  body: 'this is content',
  tags: ['gaming', 'text'],
  create_at: new Date(),
  author: authorOne,
};

let age: number = 30;
let firstName: string = 'Mario';
let isFictional: boolean;

age = 31;
firstName = 'Luigi';
isFictional = false;

let planet = 'Earth';
let moons = 1;
let isLarge = false;
planet = 'Saturn';
moons = 145;

// null & unbdefined
let something: null;
let anotherThing: undefined;

// arrays
let names: string[] = [
  'harry potter',
  'lord voldemort',
  'dumbledore',
  'snape',
  'hagrid',
  'hermione',
  'ron',
];
let ages: number[] = [25, 24, 10];
names.push('this is a test');
ages.push(1);
//object literals
const g = names[2];
console.log(g);
//type inference with object literals
let user: { firstName: string; age: number; id: number } = {
  firstName: 'hello',
  age: 30,
  id: 10,
};

console.log(user);
