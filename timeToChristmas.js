var dateChristmas = new Date("2021-12-25");
var dateNow = new Date();

var differenceInMillisecondsToChristmas = dateChristmas.getTime() - dateNow.getTime()

console.log ("Milliseconds to christmas: ", differenceInMillisecondsToChristmas);
console.log ("Seconds to christmas: ", Math.round(differenceInMillisecondsToChristmas / 1000));
console.log ("Minutes to christmas: ", Math.round(differenceInMillisecondsToChristmas/ (60*1000)));
console.log ("Hours to christmas: ", Math.round(differenceInMillisecondsToChristmas/(60*60*1000)));
console.log ("Days to christmas: ", Math.round(differenceInMillisecondsToChristmas/(60*60*24*1000)));
