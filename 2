import './App.css';
import { useEffect, useState } from 'react';

function App() {
  const [info, setInfo] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const [quoteName, setQuoteName] = useState('');
  const [quoteEmail, setQuoteEmail] = useState('');
  const [quotePhone, setQuotePhone] = useState('');
  const [quoteDetails, setQuoteDetails] = useState('');
  const [formLoading, setFormLoading] = useState(false);
  const [formResponse, setFormResponse] = useState('');
  const [formError, setFormError] = useState(null);

  useEffect(() => {
    fetch('/api/info')
      .then((res) => {
        if (!res.ok) throw new Error('Failed to fetch info');
        return res.json();
      })
      .then((data) => {
        setInfo(data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  const services = info?.services ?? [
    'Lawn mowing',
    'Grass planting',
    'Edge trimming',
    'Leaf removal',
    'Bushes & hedge shaping',
    'Flower bed care',
    'Seasonal cleanup',
  ];

  return (
    <div className="app">
      <header className="topbar">
        <h1>{info?.name ?? 'Smith/Brew Lawncare LLC'}</h1>
        <p>{info?.type ?? 'Black-owned lawncare specialists for safer, greener yards.'}</p>
      </header>

      <section className="hero">
        <h2>Beautify Your Yard</h2>
        <p>Full-service lawn care and landscaping built around trust, excellence, and local community.</p>
      </section>

      <section id="gallery" className="content-section">
        <h2>Gallery</h2>
        <div className="gallery">
          <img src="/images/lawn-1.jpg" alt="Trimmed lawn example" />
          <img src="/images/lawn-2.jpg" alt="Flower bed landscaping" />
          <img src="/images/lawn-3.jpg" alt="Clean edges and mulch" />
        </div>
        <p>Replace these with your own lawncare photos in `public/images/`.</p>
      </section>

      <section id="services" className="content-section">
        <h2>Services</h2>
        <ul>
          {services.map((service) => (
            <li key={service}>{service}</li>
          ))}
        </ul>
      </section>

      <section id="pricing" className="content-section">
        <h2>Pricing & Payment</h2>
        <p>Get a free quote for your lawn or landscape job — no commitment required.</p>
        <p>We require 30% upfront deposit to secure your appointment and ensure scheduling priority.</p>
        <ul>
          <li>Deposit: 30% at booking</li>
          <li>Balance due on completion</li>
          <li>Multiple payment methods accepted</li>
        </ul>
      </section>

      <section id="faq" className="content-section">
        <h2>FAQ</h2>
        <div className="faq">
          <details>
            <summary>How do I book a lawncare appointment?</summary>
            <p>Contact us by email or phone to schedule a service date and get a personalized quote.</p>
          </details>
          <details>
            <summary>Do you service all neighborhoods?</summary>
            <p>Yes, we serve most local neighborhoods and can discuss availability during your request.</p>
          </details>
          <details>
            <summary>What forms of payment do you accept?</summary>
            <p>We accept cash, card, and digital payments. Payment details are confirmed when booking.</p>
          </details>
        </div>
      </section>

      <section id="contact" className="content-section">
        <h2>Contact</h2>
        {loading ? (
          <p>Loading contact details...</p>
        ) : error ? (
          <p>Error: {error}</p>
        ) : (
          <>
            <p>Email: <a href={`mailto:${info.email}`}>{info.email}</a></p>
            <p>Phone: <a href={`tel:${info.phone}`}>{info.phone}</a></p>
            <p>Legal: Smith/Brew Lawncare LLC (LLC)</p>
          </>
        )}
      </section>

      <section id="request" className="content-section">
        <h2>Request a Free Quote</h2>
        <p>We require 30% upfront deposit after your quote is confirmed.</p>
        <form
          onSubmit={(e) => {
            e.preventDefault();
            setFormError(null);
            setFormResponse('');

            if (!quoteName.trim() || !quoteEmail.trim() || !quotePhone.trim() || !quoteDetails.trim()) {
              setFormError('Please fill all fields before submitting.');
              return;
            }

            if (!/^\S+@\S+\.\S+$/.test(quoteEmail)) {
              setFormError('Please enter a valid email address.');
              return;
            }

            if (!/^\+?[0-9\-\s]{7,15}$/.test(quotePhone)) {
              setFormError('Please enter a valid phone number.');
              return;
            }

            setFormLoading(true);

            fetch('/api/quotes', {
              method: 'POST',
              headers: { 'Content-Type': 'application/json' },
              body: JSON.stringify({
                name: quoteName,
                email: quoteEmail,
                phone: quotePhone,
                details: quoteDetails,
              }),
            })
              .then((res) => {
                if (!res.ok) throw new Error('Failed to send quote request');
                return res.json();
              })
              .then((data) => {
                setFormResponse('Thanks! Quote request submitted. We will contact you soon.');
                setQuoteName('');
                setQuoteEmail('');
                setQuotePhone('');
                setQuoteDetails('');
              })
              .catch((err) => {
                setFormError(err.message);
              })
              .finally(() => setFormLoading(false));
          }}
        >
          <input
            type="text"
            value={quoteName}
            onChange={(e) => setQuoteName(e.target.value)}
            placeholder="Your name"
            required
          />
          <input
            type="email"
            value={quoteEmail}
            onChange={(e) => setQuoteEmail(e.target.value)}
            placeholder="Your email"
            required
          />
          <input
            type="tel"
            value={quotePhone}
            onChange={(e) => setQuotePhone(e.target.value)}
            placeholder="Your phone"
            required
          />
          <textarea
            value={quoteDetails}
            onChange={(e) => setQuoteDetails(e.target.value)}
            placeholder="Describe your lawncare needs"
            required
          />
          <button type="submit" disabled={formLoading}>
            {formLoading ? 'Sending...' : 'Send Quote Request'}
          </button>
        </form>
        {formResponse && <p className="form-success">{formResponse}</p>}
        {formError && <p className="form-error">{formError}</p>}
      </section>

      <footer>
        <p>&copy; {new Date().getFullYear()} Smith/Brew Lawncare LLC. All rights reserved.</p>
      </footer>
    </div>
  );
}

export default App;
