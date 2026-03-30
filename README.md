# Smith/Brew Lawncare LLC

## Overview
Smith/Brew Lawncare LLC is a black-owned lawn and landscape business delivering quality lawn mowing, grass planting, edge trimming, leaf removal, bush shaping, and flower bed maintenance. We help homeowners look their best and maximize curb appeal through dependable service.

## Services Offered
- Lawn mowing
- Grass planting
- Edge trimming
- Leaf removal
- Bushes/hedge shaping
- Flower bed care
- Seasonal cleanup

## Contact Information
For inquiries or to schedule service, please reach out to us:
- Email: Marcussmith481@gmail.com
- Phone: 678-544-7911
- Legal: Smith/Brew Lawncare LLC (LLC)

## Backend + Frontend Apps Included
This repo now contains:
- `index.html`, `css/styles.css`, `js/script.js` (static site)
- `nextjs-app` (Next.js minimal page)
- `frontend/react-app` (React app + live API fetch)
- `frontend/angular-app`, `frontend/vue-app`, `frontend/svelte-app` stub guides
- `backend/ruby/app.rb` (Sinatra `/api/info`)
- `backend/php/index.php` (`/api/info` JSON)
- `backend/go/main.go` (`/api/info` JSON)
- `database/sql/schema.sql` and `database/nosql/seed.json`
- `styles/stylus`, `styles/less`, `styles/sass` examples

## Setup Instructions
1. Clone repository.
2. For static: open `index.html`.
3. For React:
  - `cd frontend/react-app && npm install && npm start`
4. For Ruby API:
  - `cd backend/ruby && bundle init && bundle add sinatra && ruby app.rb`
5. For PHP API:
  - `cd backend/php && php -S localhost:8000`
6. For Go API:
  - `cd backend/go && go run main.go`
7. Confirm frontend fetch with `http://localhost:3000` and any backend host.

## New Quote request flow
- Submit the quote form in React; it posts to `/api/quotes`.
- Go API supports `POST /api/quotes` (stores in-memory and returns an ID) and `GET /api/quotes` (list).
- Add CORS to Ruby/PHP if using PHP or Ruby backend with React on port 3000.

## Notes
- Run one backend at a time or use a proxy to avoid conflicts.
- Adjust CORS or `vite.config.js` if fetching from a different port.

- `index.html` for structure and content
- `css/styles.css` for styling and responsive layout
- `js/script.js` for any interactive behaviors

### Setup Instructions
1. Clone the repository to your local machine.
2. Ensure the `css/styles.css` file is linked correctly in `index.html`.
3. Place any required images (e.g. hero background) in an `images/` folder under the project root or adjust paths accordingly.
4. Open `index.html` in your browser to view the site and test responsiveness.
5. Customize the HTML, CSS, and JavaScript as desired for your branding or additional features.

## Additional Information
For any further questions or support, please contact us via the provided email or phone number. We look forward to making your vehicle look and feel like new!