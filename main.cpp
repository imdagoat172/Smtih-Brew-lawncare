#include <iostream>

int main() {
    std::cout << "Content-Type: text/html\n\n";
    std::cout << "<html><head><title>Smith/Brew Lawncare</title></head><body>";
    std::cout << "<h1>Smith/Brew Lawncare LLC</h1>";
    std::cout << "<p>Black-owned business: lawns, grass planting, edge trimming, leaf removal, bushes, flowers.</p>";
    std::cout << "<p>Contact: <a href='mailto:Marcussmith481@gmail.com'>Marcussmith481@gmail.com</a> | 678-544-7911</p>";
    std::cout << "<p>Legal: LLC</p>";
    std::cout << "</body></html>";
    return 0;
}
