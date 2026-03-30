<?php
header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: GET, POST, OPTIONS');
header('Access-Control-Allow-Headers: Content-Type');

if ($_SERVER['REQUEST_METHOD'] === 'OPTIONS') {
    http_response_code(200);
    exit;
}

$path = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);
$quotesFile = __DIR__ . '/quotes.json';
$quotes = file_exists($quotesFile) ? json_decode(file_get_contents($quotesFile), true) : [];

if ($path === '/api/info') {
    header('Content-Type: application/json');
    echo json_encode([
        'name' => 'Smith/Brew Lawncare LLC',
        'type' => 'Black-owned lawn care',
        'services' => ['lawn mowing', 'grass planting', 'edge trimming', 'leaf removal', 'bushes', 'flowers'],
        'email' => 'Marcussmith481@gmail.com',
        'phone' => '678-544-7911'
    ]);
    exit;
}

if ($path === '/api/quotes') {
    if ($_SERVER['REQUEST_METHOD'] === 'GET') {
        header('Content-Type: application/json');
        echo json_encode(['count' => count($quotes), 'quotes' => $quotes]);
        exit;
    }

    if ($_SERVER['REQUEST_METHOD'] === 'POST') {
        $payload = json_decode(file_get_contents('php://input'), true);
        $required = ['name', 'email', 'phone', 'details'];
        $missing = [];

        foreach ($required as $field) {
            if (empty(trim($payload[$field] ?? ''))) {
                $missing[] = $field;
            }
        }

        if ($missing) {
            http_response_code(400);
            header('Content-Type: application/json');
            echo json_encode(['error' => 'Missing fields: ' . implode(', ', $missing)]);
            exit;
        }

        $id = count($quotes) + 1;
        $quote = array_merge($payload, ['id' => $id, 'status' => 'pending']);
        $quotes[] = $quote;
        file_put_contents($quotesFile, json_encode($quotes, JSON_PRETTY_PRINT));

        http_response_code(201);
        header('Content-Type: application/json');
        echo json_encode(['id' => $id, 'message' => 'Quote request received', 'quote' => $quote]);
        exit;
    }
}

http_response_code(404);
header('Content-Type: application/json');
echo json_encode(['error' => 'Not found']);
?>
