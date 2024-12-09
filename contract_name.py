import requests
import argparse

def get_contract_name(address, api_key):
    """
    Fetches the contract name for a given Ethereum address from Etherscan.

    :param address: Ethereum contract address (str)
    :param api_key: Etherscan API key (str)
    :return: Contract name or error message
    """
    url = "https://api.etherscan.io/api"
    
    # Parameters for the Etherscan API
    params = {
        "module": "contract",
        "action": "getsourcecode",
        "address": address,
        "apikey": api_key
    }

    try:
        # Make the API request
        response = requests.get(url, params=params)
        response.raise_for_status()
        
        # Parse the response
        data = response.json()
        if data["status"] == "1" and data["message"] == "OK":
            contract_name = data["result"][0]["ContractName"]
            return contract_name if contract_name else "Contract name not found"
        else:
            return f"Error: {data['message']}"

    except requests.exceptions.RequestException as e:
        return f"Request failed: {e}"

if __name__ == "__main__":
    # Parse command-line arguments
    parser = argparse.ArgumentParser(description="Fetch Ethereum contract name from Etherscan.")
    parser.add_argument("address", help="Ethereum contract address")
    parser.add_argument(
        "--api-key", 
        required=True,
        help="Etherscan API key"
    )
    
    args = parser.parse_args()
    
    # Call the function with provided arguments
    contract_name = get_contract_name(args.address, args.api_key)
    print(f"{contract_name}")
