package mocks

import l8m "github.com/saichler/l8common/go/mocks"

var (
	firstNames = l8m.FirstNames
	lastNames  = l8m.LastNames
	cities     = l8m.Cities
)

var categoryNames = []string{
	"Electronics", "Computers & Laptops", "Smartphones", "Audio Equipment",
	"Cameras", "Furniture", "Home Appliances", "Kitchen", "Clothing",
	"Shoes", "Jewelry", "Books", "Toys & Games", "Sports Equipment",
	"Musical Instruments", "Art & Collectibles", "Tools & Hardware",
	"Garden & Outdoor", "Automotive", "Health & Beauty",
}

var listingTitles = []string{
	"Looking for a MacBook Pro 2024", "Need a standing desk", "Want quality headphones",
	"Searching for vintage vinyl records", "Need a DSLR camera", "Looking for gaming PC",
	"Want a quality espresso machine", "Need winter jacket size L", "Looking for road bike",
	"Want acoustic guitar", "Need 4K monitor", "Looking for smart watch",
	"Want quality cookware set", "Need running shoes size 10", "Looking for drone with camera",
	"Want mechanical keyboard", "Need portable speaker", "Looking for telescope",
	"Want art supplies set", "Need power drill", "Looking for yoga mat",
	"Want electric scooter", "Need tablet for drawing", "Looking for vintage watch",
	"Want quality backpack", "Need home security camera", "Looking for air purifier",
	"Want rowing machine", "Need wireless earbuds", "Looking for record player",
}

var listingDescriptions = []string{
	"Must be in excellent condition with original charger.",
	"Prefer ergonomic design with adjustable height.",
	"Looking for noise-canceling with good bass response.",
	"Classic rock or jazz, willing to pay premium for rare pressings.",
	"Full frame preferred, willing to consider mirrorless.",
	"High-end GPU required, at least 32GB RAM.",
	"Semi-automatic with built-in grinder preferred.",
	"Waterproof, insulated, brand name preferred.",
	"Carbon frame, Shimano components, size 56cm.",
	"Solid top, good for intermediate players.",
}

var bidMessages = []string{
	"I have exactly what you're looking for, barely used.",
	"Great condition, purchased last year, comes with all accessories.",
	"Selling mine as I'm upgrading, perfect condition.",
	"Like new, used only a few times, original packaging.",
	"Well maintained, no scratches or dents.",
	"Excellent condition, includes warranty transfer.",
	"Barely used, selling because I received a duplicate gift.",
	"Professional grade, meticulously cared for.",
	"Mint condition, stored properly, smoke-free home.",
	"One owner, all original parts, documentation included.",
}

var dealMessages = []string{
	"Hi, I've prepared the item for shipping. When works for you?",
	"Great, I'll ship it out tomorrow morning.",
	"Tracking number will be sent once I drop it off at the carrier.",
	"Item has been shipped! Should arrive in 3-5 business days.",
	"Thank you for the quick payment!",
	"Could you confirm the shipping address?",
	"Package was delivered, please confirm receipt.",
	"Everything looks great, thank you!",
	"Happy to do business with you!",
	"Let me know if you have any questions about the item.",
}

var reviewTitles = []string{
	"Excellent transaction!", "Smooth and easy", "Highly recommend",
	"Great seller, fast shipping", "Item as described", "Very satisfied",
	"Quick and professional", "Would buy again", "Outstanding experience",
	"Trustworthy and reliable", "Perfect condition as promised",
	"Exceeded expectations", "Fair price, great quality",
	"Communication was excellent", "Hassle-free deal",
}

var reviewContents = []string{
	"The item was exactly as described and arrived quickly. Very happy with this purchase.",
	"Smooth transaction from start to finish. The seller was communicative and professional.",
	"Great experience overall. The item was in better condition than expected.",
	"Fast shipping and the item was packaged very carefully. Would definitely recommend.",
	"Fair pricing and honest description. A pleasure to deal with.",
	"Everything went perfectly. The buyer was responsive and payment was prompt.",
	"Excellent communication throughout the process. Very trustworthy.",
	"The deal was completed quickly and efficiently. No issues at all.",
	"Item arrived in perfect condition. Very satisfied with this transaction.",
	"Would happily do business with this person again. Highly recommended.",
}

var shippingCarriers = []string{
	"USPS", "UPS", "FedEx", "DHL", "Amazon Logistics",
}
