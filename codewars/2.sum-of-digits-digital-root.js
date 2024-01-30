// https://www.codewars.com/kata/541c8630095125aba6000c00/javascript

function digitalRoot(n) {
	let res = n.toString().split('').reduce((prev, curr) => {
		return prev + +curr
	}, 0)
	if(res>=10){
		return digitalRoot(res);
	}
	return res
}

console.log(digitalRoot(16));
console.log(digitalRoot(474940));