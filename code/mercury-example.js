import Mercury from '@postlight/mercury-parser'

const useragent = 'Googlebot/2.1 (+http://www.google.com/bot.html)';
async function runMercuryParser(url) {
    data = await Mercury.parse(url, {
        headers : { 'User-Agent' : useragent },
        contentType: 'html'
    });
    return data;
}
