// https://medium.com/poka-techblog/simplify-your-javascript-use-map-reduce-and-filter-bd02c593cc2d

function isOdd(n) {
    if (n % 2 == 0) {
        return n + " :NO"
    }
    return n + " :YES"
}

async function main() {
    console.log("MAPS")
    numbers = [1,2,3,4,5,6,7,8,9,10,11] 
    console.log(numbers)
    console.log(" single line using a function ")
    console.log(numbers.map(isOdd))

    console.log("SEXY VERSION OF MAP")
    console.log(numbers.map(n => n % 2 == 0 ? "EVEN" : "ODD"))

    console.log("Boring way of doing it")
    result = []
    for (j = 0; j < numbers.length; j++) {
        result.push(isOdd(numbers[j]))
    }
    console.log(result)

    console.log("FILTERS")
    console.log(" >3  ",numbers.filter(n => n > 3))
    console.log("even ",numbers.filter(n => n % 2 == 0))

    console.log("biggest", numbers.reduce((x,n)=> x > n ? x : n))
    console.log("smallest", numbers.reduce((x,n)=> x < n ? x : n))

    console.log("COMBINING FILTER and REDUCE")

    console.log("biggest even ", numbers.filter(n => n % 2 == 0).reduce((x,n)=> x > n ? x : n))
    console.log("smallest even ", numbers.filter(n => n % 2 == 0).reduce((x,n)=> x < n ? x : n))
    console.log("biggest odd ", numbers.filter(n => n % 2 == 1).reduce((x,n)=> x > n ? x : n))
    console.log("smallest odd ", numbers.filter(n => n % 2 == 1).reduce((x,n)=> x < n ? x : n))

}


main()
    .then(()=>process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });