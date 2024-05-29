import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns

# Set the random seed for reproducibility
np.random.seed(42)

# Define the rate parameter (lambda) and the shift
lambda_param = 0.07  # Rate parameter for the exponential distribution
shift = 100  # Minimum interval time in Miliseconds

# Generate random intervals using the exponential distribution and apply the shift
num_intervals = 10000  # Number of intervals to generate
intervals = np.random.exponential(scale=1/lambda_param, size=num_intervals) + shift

# Plot the histogram of the generated intervals
plt.figure(figsize=(10, 6))
sns.histplot(intervals, bins=30, kde=True, stat="density", color='skyblue', label='Histogram of Intervals')

# Plot the theoretical PDF of the shifted exponential distribution
x = np.linspace(shift, max(intervals), 10000)
pdf = lambda_param * np.exp(-lambda_param * (x - shift))
plt.plot(x, pdf, 'r-', lw=2, label='Theoretical Shifted Exponential PDF')

# Add labels and legend
plt.xlabel('Interval Time (ms)')
plt.ylabel('Density')
plt.title('Histogram and Theoretical PDF of Shifted Exponential Distribution')
plt.legend()
plt.grid(True)

# Show the plot
plt.show()
