import os
import pandas as pd
import matplotlib.pyplot as plt

def plot_latency(log_file):
    try:
        # Read the CSV file
        data = pd.read_csv(log_file)

        # Check if required columns are present
        if 'Request' not in data.columns or 'Latency (ms)' not in data.columns:
            print(f"Skipping {log_file}: Unexpected format.")
            return

        # Sort data by request number to ensure smooth plotting
        data = data.sort_values(by=['Request'])

        # Extract request number and latency
        request_num = data['Request']
        latency = data['Latency (ms)']

        # Calculate average latency
        avg_latency = latency.mean()

        # Create the images folder if it doesn't exist
        if not os.path.exists('images'):
            os.makedirs('images')

        # Plot the latency graph
        plt.figure(figsize=(10, 6))
        plt.plot(request_num, latency, marker='o', linestyle='-', color='b', label='Latency')

        # Plot the average latency line in red
        plt.axhline(y=avg_latency, color='red', linestyle='--', label=f'Avg Latency: {avg_latency:.2f} ms')

        # Add a text label for the average latency line
        plt.text(request_num.iloc[-1], avg_latency, f'{avg_latency:.2f} ms', color='red', fontsize=12, ha='right', va='bottom')

        # Add title and labels
        plt.title(f'Latency Over Time for {log_file}')
        plt.xlabel('Request Number')
        plt.ylabel('Latency (ms)')
        plt.grid(True)
        plt.legend()

        # Save the plot to the images folder
        image_name = f"images/latency_plot_{log_file.split('.')[0]}.png"
        plt.savefig(image_name)
        plt.close()

        print(f"Latency plot saved as {image_name}")
    except Exception as e:
        print(f"Error processing {log_file}: {e}")

if __name__ == "__main__":
    # Find all log*.csv files
    log_files = [f for f in os.listdir() if f.startswith("log") and f.endswith(".csv")]
    
    # Process each CSV file
    if log_files:
        for log_file in log_files:
            plot_latency(log_file)
    else:
        print("No log files found.")
