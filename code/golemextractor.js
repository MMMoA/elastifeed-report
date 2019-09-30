export const GolemExtractor = {
    title: {
        selectors : ['article header h1'],
    },
    author: {
        selectors : ['article header .authors .authors__name'],
    },
    date_published: {
        selectors : ['article header .authors .authors__pubdate'],
    },
    dek: {
        selectors : ['article header p'],
    },
    lead_image_url: {
        selectors : ['article header figure img'],
    },
    content: {
        selectors: ['article formatted'],

        // Clean out the div containers that will have ads.
        clean: ['div'],
    },
}